import React from 'react';
import ReactDOM from 'react-dom';
import {DataTable} from 'react-data-components';
import {blueGrey500} from 'material-ui/styles/colors';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import Toggle from 'material-ui/toggle';
import { Checkbox } from 'react-bootstrap';
import { Router, Route, Link, browserHistory } from 'react-router';
import SnrTestEdit from './snrtestEdit.jsx';

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


// List all the RMA using the react-data-components.
export default class RmaList extends React.Component {
  
  constructor(props) {
    super(props);
    this.state = {
      data: {RMA:[]},
      isSelectedID: "",
    }
  }


  // At startup get all the Water Test data
  componentDidMount() {
    this.loadRmaFromServer();
    console.log("data length %i\n", this.state.data.length);
  }

  // Get the RMA data from the database using AJAX
  loadRmaFromServer() {
    $.ajax({
      url: "/vault/rma",
      dataType: 'json',
      cache: false,
      success: function(data) {
        console.log("data length %i\n", data.length);
        this.setState({data: data});
      }.bind(this),
      error: function(xhr, status, err) {
        console.error("/vault/rma", status, err.toString());
      }.bind(this)
    });
  }

    // Call API to set IsSelect selection
    apiSetSelected(selectedID) {
    var urlSelected = "/vault/rma/select/" + selectedID;
    $.ajax({
      url: urlSelected,
      dataType: 'json',
      cache: false,
      success: function(data) {
        console.log("%s %s\n", urlSelected, this.state.isSelectedID);
      }.bind(this),
      error: function(xhr, status, err) {
        console.error(urlSelected, status, err.toString());
      }.bind(this)
    });
  }

  // Convert to Bool
  convertToBool(val) {
    return (val === 'true');
  }

  // Convert bool to checked
  convertToChecked(val) {
    if(val == 'true') {
      return 'checked';
    }

    return '';
  }

  // Selection change for IsSelected Column
  handleIsSelectedChange(id) {
    // Set state
    //this.setState({isSelectedID: id});

    // Call the API
    this.apiSetSelected(id);
  }

  // Render function
  render() {

    // Report Column
    const renderReport =
      (val, row) =>
        <div>
          <Link to={"/rma/" + `${row['id']}`}> EDIT </Link>
        </div>;

    // All Columns
    var columns = [
    { title: 'RmaType', prop: 'RmaType'},
    { title: 'RmaNumber', prop: 'RmaNumber'},
    { title: 'Company', prop: 'Company'},
    { title: 'Status', prop: 'Status'},
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
        initialData={this.state.data.RMA}
        initialPageLength={20}
        initialSortBy={{ prop: 'Modified', order: 'descending' }}
        pageLengthOptions={[ 5, 20, 50 ]}
        />
      </div>
    );
  }
}

// Set the table to compTable
// Use the url PROP to get the Water Test data
//ReactDOM.render((<WaterTestCompList url="/vault/wt" selectedURL="/vault/wt/select/" editURL="watertests/" />), document.getElementById('compTable'));


