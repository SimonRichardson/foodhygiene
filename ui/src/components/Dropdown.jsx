import React from "react";
import Loader from 'react-loader';

export default class Dropdown extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            loaded: false,
            authorities: [],
            value: ''
        };


        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
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
            .catch(err => this.setState(Object.assign(this.state, { loaded: true, authorities: [] })));
    }

    handleChange(event) {
        // On change make sure that we set the state.
        this.setState(Object.assign(this.state, { value: event.target.value }));
    }

    handleSubmit(Event) {
        event.preventDefault();
    }

    render() {
        // Render the drop down along with the label.
        const submitStyle = {
            display: this.props.showSubmit ? 'block' : 'none'
        };

        return (
            <form onSubmit={this.handleSubmit}>
                <label>
                    Pick an authority to view:
                </label>
                <br />
                <Loader loaded={this.state.loaded}>
                    <select onChange={this.handleChange}>
                        {this.state.authorities.map(authority =>
                            <option value="{authority.local_id}">{authority.name}</option>
                        )}
                    </select>
                    <div style={submitStyle}>
                        <br />
                        <input type="submit" value="Submit" />
                    </div>
                </Loader>
            </form>
        )
    }
}
