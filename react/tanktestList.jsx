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


// List all the Tank Test using the react-data-components.
var TankTestList = React.createClass({
  
  // Set the initial state
  getInitialState: function() {

    // Set the STATE
    return {
      data: {TankTests:[]},
      isSelectedID: "",
    };
  },

  // At startup get all the Water Test data
  componentDidMount: function() {
    this.loadTankTestFromServer();
    console.log("data length %i\n", this.state.data.length);
  },

  // Get the Tank Test data from the database using AJAX
  loadTankTestFromServer: function() {
    $.ajax({
      url: this.props.url,
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

    // Call API to set IsSelect selection
    apiSetSelected: function(selectedID) {
    var urlSelected = this.props.selectedURL + selectedID;
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
  },

  // Convert to Bool
  convertToBool: function(val) {
    return (val === 'true');
  },

  // Selection change for IsSelected Column
  handleIsSelectedChange: function(id) {
    // Set state
    this.setState({isSelectedID: id});

    // Call the API
    this.apiSetSelected(id);
  },

  // Render function
  render: function() {

    // Report Column
    const renderReport =
      (val, row) =>
        <a href={`${row['PlotReport']}`}> Report </a>;
        

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
    { title: 'Subsystem', prop: 'SubsystemDescStr'},
    { title: 'Type', prop: 'TankTestType'},
    { title: 'Orientation', prop: 'TestOrientation'},
    { title: 'BT Distance', prop: 'BtDistance'},
    { title: 'Distance Err', prop: 'DistanceError'},
    { title: 'BT Direction', prop: 'BtDirection'},
    { title: 'Direction Err', prop: 'DirectionError'},
    { title: 'Report', render: renderReport, className: 'text-center' },
  ];

    return (
      <div>
        <DataTable
        className="container"
        keys="id"
        columns={columns}
        initialData={this.state.data.TankTests}
        initialPageLength={20}
        initialSortBy={{ prop: 'Created', order: 'descending' }}
        pageLengthOptions={[ 5, 20, 50 ]}
        />
      </div>
    );
  }
});

// Set the table to compTable
// Use the url PROP to get the Water Test data
ReactDOM.render((<TankTestList url="/vault/tt" selectedURL="/vault/tt/select/" />), document.getElementById('tanktestList'));

