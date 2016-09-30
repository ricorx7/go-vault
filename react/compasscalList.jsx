import React from 'react';
import ReactDOM from 'react-dom';
import {DataTable} from 'react-data-components';
import {blueGrey500} from 'material-ui/styles/colors';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import Toggle from 'material-ui/toggle';
import { Checkbox } from 'react-bootstrap';
import { Router, Route, Link, browserHistory } from 'react-router';
import CompassCalEdit from './compasscalEdit.jsx';

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


// List all the Compass Cal using the react-data-components.
export default class CompassCalList extends React.Component {
  
  constructor(props) {
    super(props);
    this.state = {
      data: {CompassCals:[]},
      isSelectedID: "",
    }
  }


  // At startup get all the Water Test data
  componentDidMount() {
    this.loadCompassCalFromServer();
    console.log("state data length %i\n", this.state.data.length);
  }

  // Get the Water Test data from the database using AJAX
  loadCompassCalFromServer() {
    $.ajax({
      url: "/vault/compasscal",
      dataType: 'json',
      cache: false,
      success: function(data) {
        console.log("data length %i\n", data.length);
        console.log("%v\n", data);
        this.setState({data: data});
      }.bind(this),
      error: function(xhr, status, err) {
        console.error("/vault/compasscal", status, err.toString());
      }.bind(this)
    });
  }

    // Call API to set IsSelect selection
    apiSetSelected(selectedID) {
    var urlSelected = "/vault/compasscal/select/" + selectedID;
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

      console.log("RENDER");

    // Waiting for AJAX response
    if(this.state.data == null)
    {
      console.log("Waiting for AJAX");
        return(<div>Loading...</div>);
    }
    else
    {
      console.log("render data %v\n", this.state.data);
    }

    // Report Column
    const renderReport =
      (val, row) =>
        <div>
          <Link to={"/compasscals/" + `${row['id']}`}> EDIT </Link>
        </div>;

    // IsSelected Column
    const renderIsSelected =
      (val, row) =>
        <MuiThemeProvider muiTheme={muiTheme}>
          <div style={styles.block}>
              <Toggle label="" defaultToggled={this.convertToBool(`${row['IsSelected']}`)} onToggle={this.handleIsSelectedChange.bind(this, `${row['id']}`)} style={styles.toggle} />
          </div>
        </MuiThemeProvider>;

    // All Columns
    var columns = [
    { title: 'Selected', render: renderIsSelected},
    { title: 'SerialNumber', prop: 'SerialNumber'},
    { title: 'Created', prop: 'Created'},
    { title: 'Links', render: renderReport, className: 'text-center' },
  ];

    return (
      <div>
        <DataTable
        className="container"
        keys="id"
        columns={columns}
        initialData={this.state.data.CompassCals}
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


