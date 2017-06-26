import React from 'react';
import Loader from 'react-loader';
import Dropdown from './Dropdown.jsx';
import Table from './Table.jsx';

const errorViewStyle = {
    width: '100%',
    height: '40px',
    fontSize: '2em',
    textAlign: 'center',
    color: '#ffffff',
    backgroundColor: '#f44336'
}

export default class App extends React.Component {
    constructor(props) {
        super(props);

        this.headers = new Headers();
        this.headers.set('content-type', 'application/json');

        this.state = {
            loaded: {
                select: false,
                table: false
            },
            authorities: [],
            ratings: [],
            error: ''
        };

        this.handleChange = this.handleChange.bind(this);
    }

    componentDidMount() {
        // When we mount, that's when we want to get the authorities from the
        // server. We make sure that the response is valid, if it's not then
        // we also handle that.
        fetch('/query/authorities', { headers: this.headers })
            .then(res => {
                if (res.ok) {
                    return res.json();
                }
                console.log(res);
                throw new Error('Invalid response - ' + res.statusText);
            })
            .then(res =>  this.setState(Object.assign(this.state, { authorities: res })))
            .catch(err => this.setState(Object.assign(this.state, { error: err.toString(), authorities: [] })))
            .then(res => {
                // State when everything is loaded
                const loaded = Object.assign(this.state.loaded, { select: true });
                this.setState(Object.assign(this.state, { loaded }));

                if (this.state.authorities.length > 0) {
                    const head = this.state.authorities[0];
                    this.handleChange({ localId: head.local_id.toString() });
                }
            });
    }

    handleChange(event) {
        // handleChange collects all the ratings for the local authority.
        const loaded = Object.assign(this.state.loaded, { table: false });
        this.setState(Object.assign(this.state, { loaded, error: '' }));

        fetch('/query/establishments?local_id=' + event.localId, { headers: this.headers })
            .then(res => {
                if (res.ok) {
                    return res.json();
                }
                throw new Error('Invalid response for ' + event.localId);
            })
            .then(res => this.setState(Object.assign(this.state, { ratings: res })))
            .catch(err => this.setState(Object.assign(this.state, { ratings: [] })))
            .then(res => {
                // State when everything is loaded
                const loaded = Object.assign(this.state.loaded, { table: true });
                this.setState(Object.assign(this.state, { loaded }));
            });
    }

    render() {
        // Switch on/off components depending on the state of the response.
        const errorStyle = this.state.error ? errorViewStyle : { display: 'none' };
        const viewStyle = {
            display: this.state.error ? 'none' : 'block',
            width: '100%'
        };

        // Render the various components depending on what we're loading.
        return (
            <div style={{textAlign: 'center'}}>
                <h1>Food Hygiene</h1>
                <div style={errorStyle}>{this.state.error}</div>
                <div style={viewStyle}>
                    <Loader loaded={this.state.loaded.select}>
                        <Dropdown data={this.state.authorities} onChange={this.handleChange} />
                        <Loader loaded={this.state.loaded.table}>
                            <Table data={this.state.ratings} />
                        </Loader>
                    </Loader>
                </div>
            </div>
        );
    }
}
