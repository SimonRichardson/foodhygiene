jest.dontMock('./Table.jsx');

import React from 'react';
import ReactDOM from 'react-dom';
import TestUtils from 'react-dom/test-utils';

const Table = require('./Table.jsx').default;

describe('Table', () => {

    it('should render empty', () => {
        const data = [];
        const instance = TestUtils.renderIntoDocument(<Table data={data} />);
        const options = TestUtils.scryRenderedDOMComponentsWithTag(instance, 'tr');

        expect(options.length).toEqual(0);
    });

    it('should render', () => {
        const data = [{ name: 'Bobs Burgers', rating: '3-Star'}];
        const instance = TestUtils.renderIntoDocument(<Table data={data} />);
        const options = TestUtils.scryRenderedDOMComponentsWithTag(instance, 'tr');

        // `+ 1` is to include the header
        expect(options.length).toEqual(data.length + 1);
    });

    it('should render multiple', () => {
        const data = [{ name: 'Bobs Burgers', rating: '3-Star'}, { name: 'Petes Pizza', rating: '4-Star'}];
        const instance = TestUtils.renderIntoDocument(<Table data={data} />);
        const options = TestUtils.scryRenderedDOMComponentsWithTag(instance, 'tr');

        // `+ 1` is to include the header
        expect(options.length).toEqual(data.length + 1);
    });
});
