import React from 'react';
import Dropdown from './Dropdown.jsx';

export default class App extends React.Component {
  render() {
    return (
     <div style={{textAlign: 'center'}}>
        <h1>Food Hygiene</h1>
        <Dropdown />
      </div>);
  }
}
