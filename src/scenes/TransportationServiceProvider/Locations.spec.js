import React from 'react';
import { shallow } from 'enzyme';
import { AddressElementDisplay } from 'shared/Address';
import { LocationsDisplay } from './Locations';

const defaultProps = {
  shipment: {
    delivery_address: {},
    pickup_address: {},
    has_delivery_address: false,
    has_secondary_pickup_address: false,
    secondary_pickup_address: {},
    service_member: { current_station: { address: {} } },
  },
};
describe('Locations component test', () => {
  describe('when LocationsDisplay is provided pickup and delivery address', () => {
    const shipment = {
      ...defaultProps.shipment,
      pickup_address: {
        street_address_1: '123 Disney Rd',
        city: 'Los Angeles',
        state: 'CA',
        postal_code: '90210',
      },
      has_delivery_address: true,
      delivery_address: {
        street_address_1: '560 State Street',
        city: 'New York',
        state: 'NY',
        postal_code: '094321',
      },
    };
    const wrapper = shallow(<LocationsDisplay shipment={shipment} />);
    it('should render 2 headers', () => {
      const headers = wrapper.find('.column-subhead');
      expect(headers.length).toBe(2);
    });
    it('should render the Pickup header', () => {
      const Pickup = wrapper.find('.column-subhead').getElements()[0];
      const { className, children } = Pickup.props;
      expect(className).toBe('column-subhead');
      expect(children).toBe('Pickup');
    });
    it('should render the Delivery header', () => {
      const Delivery = wrapper.find('.column-subhead').getElements()[1];
      const { className, children } = Delivery.props;
      expect(className).toBe('column-subhead');
      expect(children).toBe('Delivery');
    });
    it('should render 2 AddressElementDisplays', () => {
      const AddressElement = wrapper.find(AddressElementDisplay);
      expect(AddressElement.length).toBe(2);
    });
  });
  describe('when LocationsDisplay is provided pickup and no delivery address it defaults to duty station', () => {
    const shipment = {
      ...defaultProps.shipment,
      pickup_address: {
        street_address_1: '123 Disney Rd',
        city: 'Los Angeles',
        state: 'CA',
        postal_code: '90210',
      },
      service_member: {
        current_station: {
          address: {
            city: 'San Diego',
            state: 'CA',
            postal_code: '92104',
          },
          name: 'Fort Worth',
        },
      },
    };
    const wrapper = shallow(<LocationsDisplay shipment={shipment} />);

    it('should render 2 AddressElementDisplays', () => {
      const AddressElement = wrapper.find(AddressElementDisplay);
      expect(AddressElement.length).toBe(2);
    });
  });
});
