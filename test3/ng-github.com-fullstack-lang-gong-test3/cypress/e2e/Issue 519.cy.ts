describe('Issue 519.cy.ts', () => {
  it('Test check box', () => {
    cy.visit('http://localhost:4200');

    // Locate the checkbox element by CSS selector, id, or another suitable selector
    cy.get('#checkbox-Default-input')  // Replace #checkbox-id with the actual selector of your checkbox
      .check()               // Clicks or checks the checkbox
      .should('be.checked'); // Assert that the checkbox is now checked
  })
})