import React from "react";

export default class Dropdown extends React.Component {
    constructor(props) {
        super(props);

        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleChange(event) {
        this.props.onChange({ localId: event.target.value });
    }

    handleSubmit(event) {
        event.preventDefault();

        this.props.onChange({ localId: event.target.value });
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
                <select onChange={this.handleChange}>
                    {
                        this.props.data.map((authority, i) =>
                            <option key={i} value={authority.local_id}>{authority.name}</option>
                        )
                    }
                </select>
                <div style={submitStyle}>
                    <br />
                    <input type="submit" value="Submit" />
                </div>
            </form>
        );
    }
}
