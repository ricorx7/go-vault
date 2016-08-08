import React from 'react';
import ReactDOM from 'react-dom';
import {DataTable} from 'react-data-components';
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
export default class WaterTestEdit extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
        data: {}
    }
  }

    // At startup get all the Water Test data
  componentDidMount() {
    this.apiGetWtSelected();
    console.log("data length %i\n", this.state.data);
  }

    // Call API to set IsSelect selection
    apiGetWtSelected() {
    var urlSelected = "/vault/wt/edit/" + this.props.params.id;
    $.ajax({
      url: urlSelected,
      dataType: 'json',
      cache: false,
      success: function(data) {
        console.log("Data gotten from %s\n", urlSelected);
        console.log("%v\n", data);
        this.setState({data: data});
      }.bind(this),
      error: function(xhr, status, err) {
        console.error(urlSelected, status, err.toString());
      }.bind(this)
    });
  }


  // Render function
  render() {
    return (
        <div>
        <h1>{this.state.data.SerialNumber}</h1>
        </div>
    );
  }
}

// Use the url PROP to get the Water Test data
//ReactDOM.render((<WaterTestEdit url="/vault/wt/edit/" />), document.getElementById('watertestEdit'));
