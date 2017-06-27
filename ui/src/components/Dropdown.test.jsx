jest.dontMock('./Dropdown.jsx');

import React from 'react';
import ReactDOM from 'react-dom';
import TestUtils from 'react-dom/test-utils';

const Dropdown = require('./Dropdown.jsx').default;

describe('Dropdown', () => {

    it('should render empty', () => {
        const data = [];
        const instance = TestUtils.renderIntoDocument(<Dropdown data={data} />);
        const options = TestUtils.scryRenderedDOMComponentsWithTag(instance, 'option');

        expect(options.length).toEqual(0);
    });

    it('should render', () => {
        const data = [{ name: 'Yorkshire', local_id: '123'}];
        const instance = TestUtils.renderIntoDocument(<Dropdown data={data} />);
        const options = TestUtils.scryRenderedDOMComponentsWithTag(instance, 'option');

        expect(options.length).toEqual(data.length);
    });

    it('should render multiple', () => {
        const data = [{ name: 'Yorkshire', local_id: '123'}, {name: 'London', local_id: '666'}];
        const instance = TestUtils.renderIntoDocument(<Dropdown data={data} />);
        const options = TestUtils.scryRenderedDOMComponentsWithTag(instance, 'option');

        expect(options.length).toEqual(data.length);
    });
});
