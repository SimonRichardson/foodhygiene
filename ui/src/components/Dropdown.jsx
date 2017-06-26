import React from "react";

export default class Dropdown extends React.Component {
    constructor(props) {
        super(props);

        this.handleChange = this.handleChange.bind(this);
    }

    handleChange(event) {
        this.props.onChange({ localId: event.target.value });
    }

    render() {
        return (
            <form>
                <label>
                    Pick an authority to view:
                </label>
                <br />
                <select onChange={this.handleChange} style={{margin: '10px 0 40px 0'}}>
                    {
                        this.props.data.map((authority, i) =>
                            <option key={i} value={authority.local_id}>{authority.name}</option>
                        )
                    }
                </select>
            </form>
        );
    }
}
