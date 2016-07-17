const React = require('react');
const ReactDOM = require('react-dom');
var DataTable = require('react-data-components').DataTable;
import {blueGrey500} from 'material-ui/styles/colors';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import Toggle from 'material-ui/toggle';

const muiTheme = getMuiTheme({
  palette: {
    accent1Color: blueGrey500,
  },
});

const style = {
  margin: 12,
};

const styles = {
  block: {
    maxWidth: 150,
  },
  toggle: {
    marginBottom: 10,
  },
};

const renderMapUrl =
  (val, row) =>
    <a href={`${row['PlotReport']}`}> Report </a>;


function handleIsSelectedChange(wt) {
    // Use API to change the value
    console.log(wt.id)
  };

const renderIsSelected =
  (val, row) =>
              <MuiThemeProvider muiTheme={muiTheme}>
                <div style={styles.block}>
                    <Toggle label="" defaultToggled={`${row['IsSelected']}`}  style={styles.toggle} />
                </div>
              </MuiThemeProvider>;

var columns = [
  { title: 'Selected', render: renderIsSelected},
  { title: 'SerialNumber', prop: 'SerialNumber'},
  { title: 'Subsystem', prop: 'SubsystemDescStr'},
  { title: 'Orientation', prop: 'TestOrientation'},
  { title: 'GpsDistance', prop: 'GpsDistance'},
  { title: 'BT Distance', prop: 'BtDistance'},
  { title: 'Distance Err', prop: 'DistanceError'},
  { title: 'GpsDirection', prop: 'GpsDirection'},
  { title: 'BT Direction', prop: 'BtDirection'},
  { title: 'Direction Err', prop: 'DirectionError'},
  { title: 'Report', render: renderMapUrl, className: 'text-center' },
];


var WaterTestCompList = React.createClass({
  getInitialState: function() {

    console.log("constructor");
    this.loadWaterTestFromServer();
    
    return {
      filter: "",
      data: {WaterTests:[]},
    };
  },

  componentDidMount: function() {
    this.loadWaterTestFromServer();
    console.log("data length %i\n", this.state.data.length);
  },

  loadWaterTestFromServer: function() {
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

  render: function() {
    return (
      <div>
<DataTable
      className="container"
      keys="id"
      columns={columns}
      initialData={this.state.data.WaterTests}
      initialPageLength={20}
      initialSortBy={{ prop: 'Created', order: 'descending' }}
      pageLengthOptions={[ 5, 20, 50 ]}
    />
      </div>
    );
  }
});

ReactDOM.render((<WaterTestCompList url="/vault/wt" />), document.getElementById('compTable'));

