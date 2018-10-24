import { testPremoveSurvey } from '../../support/testPremoveSurvey';

/* global cy */
describe('TSP User Ships a Shipment', function() {
  beforeEach(() => {
    cy.signIntoTSP();
  });
  it('tsp user Picks Up a shipment', function() {
    tspUserPacksShipment();
    tspUserPicksUpShipment();
    tspUserDeliversShipment();
  });

  it('tsp user enters a delivery date', function() {
    tspUserVisitsAnInTransitShipment('ENTDEL');
    tspUserVerifiesShipmentStatus('In_transit');
    tspUserCancelsEnteringADeliveryDate();
    tspUserEntersADeliveryDate();
    tspUserVerifiesShipmentStatus('Delivered');
  });
});

function tspUserVerifiesShipmentStatus(status) {
  cy.get('li').contains(`Status: ${status}`);
}

function tspUserCancelsEnteringADeliveryDate() {
  cy
    .get('button')
    .contains('Enter Delivery')
    .should('exist');

  cy
    .get('button')
    .contains('Enter Delivery')
    .click();

  cy
    .get('input[name="actual_delivery_date"]')
    .type('10/24/2018')
    .blur();

  cy
    .get('a')
    .contains('Cancel')
    .click();

  cy.get('input[name="actual_delivery_date"]').should('not.exist');
}

function tspUserEntersADeliveryDate() {
  cy
    .get('button')
    .contains('Enter Delivery')
    .click();

  cy.get('input[name="actual_delivery_date"]').should('be.empty');

  cy
    .get('input[name="actual_delivery_date"]')
    .type('10/24/2018')
    .blur();

  cy
    .get('button')
    .contains('Done')
    .click();

  cy.get('button').should('not.contain', 'Enter Delivery');
}

function tspUserVisitsAnInTransitShipment(locator) {
  cy.visit('/queues/in_transit');

  cy.location().should(loc => {
    expect(loc.pathname).to.match(/^\/queues\/in_transit/);
  });

  cy
    .get('div')
    .contains(locator)
    .dblclick();
}

function tspUserPacksShipment() {
  // Open approved shipments queue
  cy
    .get('div')
    .contains('Approved Shipments')
    .click();

  cy.location().should(loc => {
    expect(loc.pathname).to.match(/^\/queues\/approved/);
  });

  // Find shipment and open it
  cy
    .get('div')
    .contains('SHIPME')
    .dblclick();

  // Click the Pack button
  cy
    .get('div')
    .contains('Enter Packing')
    .click();

  // Done button should be disabled.
  cy
    .get('button')
    .contains('Done')
    .should('be.disabled');

  // Pick a date!
  cy
    .get('div')
    .contains('Actual Pack Date')
    .get('input')
    .click();

  cy
    .get('div.DayPicker-Month')
    .contains('10')
    .click();

  // Cancel
  cy
    .get('button')
    .contains('Cancel')
    .click();

  // Check that the date doesn't appear in dates panel
  cy.get('div.actual_pack_date').contains('missing');

  // Wash, Rinse, Repeat
  // Click the Pack button
  cy
    .get('div')
    .contains('Enter Packing')
    .click();

  // Done button should be disabled.
  cy
    .get('button')
    .contains('Done')
    .should('be.disabled');

  // Pick a date!
  cy
    .get('div')
    .contains('Actual Pack Date')
    .get('input')
    .click();

  cy
    .get('div.DayPicker-Day')
    .contains('10')
    .click();

  cy
    .get('button')
    .contains('Done')
    .click();

  // Appears in dates panel
  cy.get('div.actual_pack_date').contains('10');
}
function tspUserPicksUpShipment() {
  cy.location().should(loc => {
    expect(loc.pathname).to.match(/^\/shipments\/[^/]+/);
  });

  // Click the Transport button
  cy
    .get('div')
    .contains('Enter Pickup')
    .click();

  // Done button should be disabled.
  cy
    .get('button')
    .contains('Done')
    .should('be.disabled');

  // Pick a date!
  cy
    .get('div')
    .contains('Actual Pickup Date')
    .get('input')
    .click();

  cy
    .get('div')
    .contains('11')
    .click();

  // Cancel
  cy
    .get('button')
    .contains('Cancel')
    .click();

  // Wash, Rinse, Repeat
  // Click the Transport button
  cy
    .get('div')
    .contains('Enter Pickup')
    .click();

  // Done button should be disabled.
  cy
    .get('button')
    .contains('Done')
    .should('be.disabled');

  // Pick a date!
  cy
    .get('div')
    .contains('Actual Pickup Date')
    .get('input')
    .click();

  cy
    .get('div')
    .contains('11')
    .click();

  cy
    .get('button')
    .contains('Done')
    .click();

  // Appears in dates panel
  cy.get('div.actual_pickup_date').contains('11');

  // New status
  cy.get('li').contains('In_transit');
}

function tspUserDeliversShipment() {
  cy.location().should(loc => {
    expect(loc.pathname).to.match(/^\/shipments\/[^/]+/);
  });

  // Click the Transport button
  cy
    .get('div')
    .contains('Enter Delivery')
    .click();

  // Done button should be disabled.
  cy
    .get('button')
    .contains('Done')
    .should('be.disabled');

  // Pick a date!
  cy
    .get('div')
    .contains('Actual Delivery Date')
    .get('input')
    .click();

  cy
    .get('div')
    .contains('13')
    .click();

  // Cancel
  cy
    .get('a')
    .contains('Cancel')
    .click();

  // Wash, Rinse, Repeat
  // Click the Transport button
  cy
    .get('div')
    .contains('Enter Delivery')
    .click();

  // Done button should be disabled.
  cy
    .get('button')
    .contains('Done')
    .should('be.disabled');

  // Pick a date!
  cy
    .get('div')
    .contains('Actual Delivery Date')
    .get('input')
    .click();

  cy
    .get('div')
    .contains('13')
    .click();

  cy
    .get('button')
    .contains('Done')
    .click();

  // New status
  cy.get('li').contains('Delivered');
}
