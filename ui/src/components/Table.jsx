import React from "react";

export default class Table extends React.Component {
    render() {
        if (this.props.data.length < 1) {
            return (<div></div>);
        }

        return (
            <table style={{textAlign: 'left', width: '400px', margin: 'auto', border: '1px solid #c1c1c1'}}>
                <thead>
                    <tr>
                        <th>Rating</th>
                        <th style={{textAlign: 'right'}}>Percentage</th>
                    </tr>
                </thead>
                <tbody>
                    {
                        this.props.data.map((rating, i) =>
                            <tr key={i}>
                                <td>{rating.name}</td>
                                <td style={{textAlign: 'right'}}>{rating.rating}</td>
                            </tr>
                        )
                    }
                </tbody>
            </table>
        );
    }
}
