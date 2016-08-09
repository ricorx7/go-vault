import React from 'react';
import ReactDOM from 'react-dom';
import {DataTable} from 'react-data-components';
import {blueGrey500} from 'material-ui/styles/colors';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import Toggle from 'material-ui/toggle';
import { Checkbox } from 'react-bootstrap';
import { Router, Route, Link, browserHistory } from 'react-router';
import WaterTestEdit from './watertestEdit.jsx';

// Theme for material-ui toggle
const muiTheme = getMuiTheme({
  palette: {
    accent1Color: blueGrey500,
  },
});

// Style for material-ui toggle
const styles = {
  block: {
    maxWidth: 150,
  },
  toggle: {
    marginBottom: 10,
  },
};


// List all the ADCP using the react-data-components.
export default class AdcpList extends React.Component {
  
  constructor(props) {
    super(props);
    this.state = {
      data: {Adcps:[]},
    }
  }


  // At startup get all the ADCP data
  componentDidMount() {
    this.loadFromServer();
    console.log("data length %i\n", this.state.data.length);
  }

  // Get the ADCP data from the database using AJAX
  loadFromServer() {
    $.ajax({
      url: "/vault/adcp",
      dataType: 'json',
      cache: false,
      success: function(data) {
        console.log("data length %i\n", data.length);
        this.setState({data: data});
      }.bind(this),
      error: function(xhr, status, err) {
        console.error("/vault/adcp", status, err.toString());
      }.bind(this)
    });
  }

  // Render function
  render() {

    // Report Column
    const renderReport =
      (val, row) =>
        <div>
          <Link to={"/adcps/" + `${row['id']}`}> EDIT </Link>
          <Link to={"/adcps/cert/" + `${row['SerialNumber']}`}> CERT </Link>
          <Link to={"/adcps/test/" + `${row['id']}`}> TEST </Link>
        </div>;

    // All Columns
    var columns = [
    { title: 'SerialNumber', prop: 'SerialNumber'},
    { title: 'Customer', prop: 'Customer'},
    { title: 'Created', prop: 'Created'},
    { title: 'Modified', prop: 'Modified'},
    { title: 'Links', render: renderReport, className: 'text-center' },
  ];

    return (
      <div>
        <DataTable
        className="container"
        keys="id"
        columns={columns}
        initialData={this.state.data.Adcps}
        initialPageLength={20}
        initialSortBy={{ prop: 'Created', order: 'descending' }}
        pageLengthOptions={[ 5, 20, 50 ]}
        />
      </div>
    );
  }
}

// Set the table to compTable
// Use the url PROP to get the Water Test data
//ReactDOM.render((<WaterTestCompList url="/vault/wt" selectedURL="/vault/wt/select/" editURL="watertests/" />), document.getElementById('compTable'));


