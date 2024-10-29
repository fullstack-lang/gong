describe('template spec', () => {
  it('Test delete', () => {
    cy.visit('http://localhost:4200');

    // Set zoom level to 50%
    cy.document().then((doc) => {
      doc.body.style.zoom = '0.5';
    });

    cy.get('#nodeA\\ \\(5\\)').click();

    cy.wait(500);

    cy.get('#Delete\\ Icon\\ 5').click();
  })
})