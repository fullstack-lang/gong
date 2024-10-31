describe('Issue 519.cy.ts', () => {
  it('Test check box', () => {
    cy.visit('http://localhost:4200');

    // Set zoom level to 50%
    cy.document().then((doc) => {
      doc.body.style.zoom = '0.7';
    });

    // Locate the checkbox element by CSS selector, id, or another suitable selector
    cy.get('#checkbox-Default-input')  // select "Default" diagram
      .check(); // Clicks or checks the checkbox

    cy.wait(500);
    cy.get('#checkbox-Default-input').should('be.checked'); // Assert that the checkbox is now checked

    cy.get('#button-Default\\ draw').click(); // Click to edit diagram Default

    cy.wait(500);

    cy.get('#checkbox-A-input').click(); // Click to add the "A" struct

  })
})