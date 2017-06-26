import React from 'react';
import Loader from 'react-loader';
import Dropdown from './Dropdown.jsx';
import Table from './Table.jsx';

export default class App extends React.Component {
    constructor(props) {
        super(props);

         this.state = {
            loaded: false,
            authorities: [],
            ratings: []
         };

        this.handleChange = this.handleChange.bind(this);
    }

    componentDidMount() {
        // When we mount, that's when we want to get the authorities from the
        // server. We make sure that the response is valid, if it's not then
        // we also handle that.
        const headers = new Headers();
        headers.set('content-type', 'application/json');

        const init = { headers };
        fetch(`/query/authorities`, init)
            .then(res => {
                if (res.ok) {
                    return res.json();
                }
                throw new Error('Invalid response');
            })
            .then(res => this.setState(Object.assign(this.state, { loaded: true, authorities: res })))
            .catch(err => this.setState(Object.assign(this.state, { loaded: true, authorities: [] })))
            .then(res => {
                // State when everything is loaded
                if (this.state.authorities.length > 0) {
                    const head = this.state.authorities[0];
                    this.handleChange({ localId: head.local_id.toString() });
                }
            });
    }

    handleChange(event) {
        console.log(event);
    }

    render() {
        return (
            <div style={{textAlign: 'center'}}>
                <h1>Food Hygiene</h1>
                <Loader loaded={this.state.loaded}>
                    <Dropdown data={this.state.authorities} onChange={this.handleChange} />
                    <Table data={this.state.ratings} />
                </Loader>
            </div>
        );
    }
}
