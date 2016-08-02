const React = require('react');
const ReactDOM = require('react-dom');
var DataTable = require('react-data-components').DataTable;
import {blueGrey500} from 'material-ui/styles/colors';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import Toggle from 'material-ui/toggle';

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


// Edit the water test data.
var WaterTestEdit = React.createClass({
  
  // Set the initial state
  getInitialState: function() {

    // Set the STATE
    return {
      data: {}
    }
  },

  // At startup get all the Water Test data
  componentDidMount: function() {
    this.loadWaterTestFromServer(this.props.id);
    console.log("data length %i\n", this.state.data.length);
  },

  // Get the Water Test data from the database using AJAX
  loadWaterTestFromServer: function(selectedID) {
      var urlID = this.props.url + selectedID;
    $.ajax({
      url: urlID,
      dataType: 'json',
      cache: false,
      success: function(data) {
        console.log("data length %i\n", data.length);
        this.setState({data: data});
      }.bind(this),
      error: function(xhr, status, err) {
        console.error(this.props.url, status, err.toString());
      }.bind(this)
    });
  },


  // Render function
  render: function() {


    return (
        <div>
        {this.state.data.SerialNumber}
        </div>
    );
  }
});

// Set the table to compTable
// Use the url PROP to get the Water Test data
ReactDOM.render((<WaterTestEdit url="/vault/wt/edit/" />), document.getElementById('watertestEdit'));

