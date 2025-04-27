package models

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

// GrabGeneratedSVGFile waits for the SVG text content to be updated via the proxy,
// receives the text through a channel, writes it to a file, and handles timeouts/errors.
// It returns the svg instance processed and an error if any step failed.
func GrabGeneratedSVGFile(stage *Stage, imageFilePath string, timeout time.Duration) (*SVG, error) {

	var svgInstance *SVG
	// Find the existing SVG instance
	for inst := range *GetGongstructInstancesSet[SVG](stage) {
		svgInstance = inst
		break // Assuming the first one is the target
	}
	// Handle case where no SVG instance exists
	if svgInstance == nil {
		return nil, errors.New("no SVG instance found in the stage")
	}

	var svgTextInstance *SvgText
	// Find the existing SvgText instance
	for inst := range *GetGongstructInstancesSet[SvgText](stage) {
		svgTextInstance = inst
		break // Assuming the first one is the target
	}

	// If no SvgText exists, create a new one
	if svgTextInstance == nil {
		svgTextInstance = new(SvgText).Stage(stage)
		// log.Println("Created new svg.SvgText instance.")
	}

	// Ensure the proxy is cleaned up if it was temporarily set
	// And reset the generation flag regardless of success/failure path
	defer func() {
		if svgInstance != nil { // Check svgInstance isn't nil before accessing
			svgInstance.IsSVGFileGenerated = false
		}
		// Only clear the proxy if we temporarily set it for this operation.
		// This assumes the proxy is *only* for this function's purpose.
		// If the proxy has a longer lifecycle, this cleanup might be incorrect.
		if svgTextInstance != nil {
			// If we just created the SvgText instance, we might want to remove it
			// Or if we are just setting the proxy temporarily:
			// svgTextInstance.Proxy = nil // Be careful with this assumption
		}

		// Commit the final state (flag reset, potentially proxy cleared)
		// Consider error handling for Commit if necessary
		stage.Commit()
		log.Println("SVG generation flag reset and final state committed.")

		// to avoid successive svg commit
		time.Sleep(500 * time.Millisecond)
	}()

	// Create a channel to receive the SvgText content
	updateChan := make(chan string, 1) // Buffer of 1 might prevent blocking sender if receiver is slow

	svgInstance.IsSVGFileGenerated = true

	// Assign the proxy with the channel
	// Ensure the SvgTextProxy implementation properly handles sending
	// data and potentially closing the channel.
	svgTextInstance.Proxy = &SvgTextProxy{
		stage:      stage,
		svgText:    svgTextInstance,
		updateChan: updateChan, // Pass the channel to the proxy
	}

	// Commit the stage to trigger the frontend update and proxy interaction.
	// Consider potential errors from Commit() if it can return them.
	stage.Commit() // Assuming this triggers the async update

	var svgContent string
	var err error // Declare error variable for the scope

	// Wait for the SvgTextProxy.Updated() method to send the SVG text content OR timeout
	log.Println("Waiting for SVG text content via channel...")
	select {
	case receivedContent, ok := <-updateChan:
		if !ok {
			// Channel was closed without sending data (unexpected if proxy should always send)
			log.Println("Warning: SVG update channel closed unexpectedly.")
			err = errors.New("SVG update channel closed without providing content")
			// Reset flag and commit is handled by defer
			return svgInstance, err
		}
		svgContent = receivedContent
		log.Println("SVG text content received.")

	case <-time.After(timeout):
		log.Printf("Error: Timed out after %v waiting for SVG text content.\n", timeout)
		err = fmt.Errorf("timed out waiting for SVG content after %v", timeout)
		// Reset flag and commit is handled by defer.
		// We might need to explicitly clean up the proxy or signal cancellation here.
		// Potentially set svgTextInstance.Proxy = nil here if it's temporary?
		return svgInstance, err
	}

	// Close the channel (if the sender doesn't). Best practice: sender closes.
	// close(updateChan) // Typically sender closes, confirm SvgTextProxy.Updated does this.

	// Now, write the received SVG content to the file
	if svgContent == "" {
		log.Println("Warning: Received empty SVG content from proxy. File will not be written.")
		// Decide if empty content is an error or acceptable
		// return svgInstance, errors.New("received empty SVG content") // Uncomment if empty is an error
	} else {
		// Use 0644 for standard file permissions
		writeErr := os.WriteFile(imageFilePath, []byte(svgContent), 0644)
		if writeErr != nil {
			log.Printf("Error writing SVG file '%s': %v\n", imageFilePath, writeErr)
			// Return the write error. The defer will handle cleanup.
			return svgInstance, fmt.Errorf("failed to write SVG file '%s': %w", imageFilePath, writeErr)
		}
		log.Printf("SVG file created successfully: %s\n", imageFilePath)
	}

	// Cleanup (resetting flag, final commit) is handled by the defer statement.
	// Return the svg instance and nil error on success
	return svgInstance, nil
}

// SvgTextProxy implements the SvgTextProxyInterface to intercept updates
// to the SvgText field and send the content through a channel.
type SvgTextProxy struct {
	stage      *Stage
	svgText    *SvgText      // Reference to the SvgText instance being proxied
	updateChan chan<- string // Use send-only channel type annotation
}

// Updated is the method called by the gong framework when the SvgText instance might have changed.
// It now matches the SvgTextProxyInterface signature: Updated().
func (p *SvgTextProxy) Updated() {
	// Check if the proxy and its fields are valid
	if p == nil || p.svgText == nil || p.updateChan == nil {
		log.Println("SvgTextProxy.Updated: Proxy or its fields are nil, skipping.")
		return
	}
	// get the value that has been updated
	p.stage.Checkout()

	// Get the current text value from the proxied SvgText instance
	textValue := p.svgText.Text // Read the value directly

	log.Printf("SvgTextProxy: Updated() called. Current SvgText.Text length: %d. Sending to channel.", len(textValue))

	// Send the content through the channel.
	// Use a select with a default case to prevent blocking if the channel is full
	// or if the receiver has already timed out and is no longer listening.
	select {
	case p.updateChan <- textValue:
		log.Println("SvgTextProxy: Content sent to channel.")
	default:
		// This case handles situations where the channel send would block.
		// This might happen if the channel buffer is full (unlikely with buffer 1 unless called rapidly)
		// or if the receiver in grabGeneratedSVGFile has already timed out and moved on.
		log.Println("SvgTextProxy: Warning - Could not send text to channel (possibly buffer full or receiver timed out).")
	}

	// IMPORTANT: The sender should close the channel when it's done sending.
	// Assuming only one update is expected per proxy instance lifecycle for this operation.
	// Ensure this doesn't cause a "send on closed channel" panic if Updated can be called multiple times.
	// Consider adding logic (e.g., a flag within the proxy) to ensure close(p.updateChan) is called only once.
	// For now, we assume Updated is called effectively once for the desired value.
	close(p.updateChan)
	log.Println("SvgTextProxy: Channel closed.")

}
