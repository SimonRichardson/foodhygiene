import React from "react";

export default class Table extends React.Component {
    render() {
        if (this.props.data.length < 1) {
            return (<div></div>);
        }

        return (
            <table>
                <thead>
                    <tr>
                        <th>Rating</th>
                        <th>Percentage</th>
                    </tr>
                </thead>
                <tbody>
                    {
                        this.props.data.map((rating, i) =>
                            <tr key={i}>
                                <td>{rating.name}</td>
                                <td>{rating.rating}</td>
                            </tr>
                        )
                    }
                </tbody>
            </table>
        );
    }
}
