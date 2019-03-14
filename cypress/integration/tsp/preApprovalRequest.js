import {
  fillAndSavePreApprovalRequest,
  editPreApprovalRequest,
  deletePreApprovalRequest,
} from '../../support/preapprovals/testCreateRequest';
import { addLegacyRequest, add105, add35A } from '../../support/preapprovals/testRobustAccessorials';

/* global cy */
describe('TSP user interacts with pre approval request panel', function() {
  beforeEach(() => {
    cy.signIntoTSP();
  });
  it('TSP user creates pre approval request', function() {
    tspUserCreatesPreApprovalRequest();
  });
  it('TSP user edits pre approval request', function() {
    tspUserEditsPreApprovalRequest();
  });
  it('TSP user deletes pre approval request', function() {
    tspUserDeletesPreApprovalRequest();
  });
  it('Add legacy 105B/E to verify they display correctly', function() {
    test105beLegacy();
  });
  it('TSP user creates 105B/E request', function() {
    test105be();
  });
  it('Add legacy 35A to verify it displays correctly', function() {
    test35ALegacy();
  });
  it('TSP user creates 35A request', function() {
    test35A();
  });
});

function tspUserCreatesPreApprovalRequest() {
  // Open new shipments queue
  cy.location().should(loc => {
    expect(loc.pathname).to.match(/^\/queues\/new/);
  });

  // Find shipment and open it
  cy.selectQueueItemMoveLocator('DATESP');

  cy.location().should(loc => {
    expect(loc.pathname).to.match(/^\/shipments\/[^/]+/);
  });

  // TSP users should not able to approve pre approval requests
  cy.get('[data-test=approve-request]').should('not.exist');

  fillAndSavePreApprovalRequest();
  // Verify data has been saved in the UI
  cy.get('tr[data-cy="130B"]').should(td => {
    const text = td.text();
    expect(text).to.include('Bulky Article: Motorcycle/Rec vehicle');
  });
}

function tspUserEditsPreApprovalRequest() {
  // Open new shipments queue
  cy.location().should(loc => {
    expect(loc.pathname).to.match(/^\/queues\/new/);
  });

  // Find shipment and open it
  cy.selectQueueItemMoveLocator('DATESP');

  cy.location().should(loc => {
    expect(loc.pathname).to.match(/^\/shipments\/[^/]+/);
  });

  editPreApprovalRequest();
  // Verify data has been saved in the UI
  cy.get('tr[data-cy="130B"]').should(td => {
    const text = td.text();
    expect(text).to.include('edited');
  });
}

function tspUserDeletesPreApprovalRequest() {
  // Open new shipments queue
  cy.location().should(loc => {
    expect(loc.pathname).to.match(/^\/queues\/new/);
  });

  // Find shipment and open it
  cy.selectQueueItemMoveLocator('DATESP');

  cy.location().should(loc => {
    expect(loc.pathname).to.match(/^\/shipments\/[^/]+/);
  });

  deletePreApprovalRequest();
  cy
    .get('.pre-approval-panel td')
    .first()
    .should('not.contain', 'Bulky Article: Motorcycle/Rec vehicle');
}

function test105beLegacy() {
  cy.selectQueueItemMoveLocator('DATESP');

  addLegacyRequest({ code: '105B', quantity1: 12 }).then(res => {
    expect(res.status).to.equal(201);
  });

  addLegacyRequest({ code: '105E', quantity1: 90 }).then(res => {
    expect(res.status).to.equal(201);
  });

  // must reload page because original 105B/E are added by cy.request()
  cy.reload();
  cy.get('td[details-cy="105B-default-details"]').should('contain', '12.0000 notes notes 105B');
  cy.get('td[details-cy="105E-default-details"]').should('contain', '90.0000 notes notes 105E');
}

function test105be() {
  cy.selectQueueItemMoveLocator('DATESP');

  add105({ code: '105B', itemSize: 30, crateSize: 25 });
  cy
    .get('td[details-cy="105B-details"]')
    .should(
      'contain',
      'description description 105B Crate: 25" x 25" x 25" (9.04 cu ft) Item: 30" x 30" x 30" notes notes',
    );

  add105({ code: '105E', itemSize: 40, crateSize: 50 });
  cy
    .get('td[details-cy="105E-details"]')
    .should(
      'contain',
      'description description 105E Crate: 50" x 50" x 50" (72.33 cu ft) Item: 40" x 40" x 40" notes notes',
    );
}

function test35ALegacy() {
  cy.selectQueueItemMoveLocator('DATESP');

  addLegacyRequest({ code: '35A', quantity1: 12 }).then(res => {
    expect(res.status).to.equal(201);
  });

  // must reload page because original 105B/E are added by cy.request()
  cy.reload();
  cy.get('td[details-cy="105B-default-details"]').should('contain', '12.0000 notes notes 105B');
  cy.get('td[details-cy="105E-default-details"]').should('contain', '90.0000 notes notes 105E');
}

function test35A() {
  cy.selectQueueItemMoveLocator('DATESP');

  add35A({});
  cy
    .get('td[details-cy="35A-details"]')
    .should('contain', 'description description 35A reason reason 35A Est. not to exceed: $250.00 Actual Amount: --');
}
