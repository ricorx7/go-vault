import React from 'react';
import ReactDOM from 'react-dom';
import {deepOrange500} from 'material-ui/styles/colors';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import Toggle from 'material-ui/toggle';

const muiTheme = getMuiTheme({
  palette: {
    accent1Color: deepOrange500,
  },
});

class World extends React.Component {
  render() {
  const styles = {
    block: {
      maxWidth: 250,
    },
    toggle: {
      marginBottom: 16,
    },
  };


    return <div>
  <MuiThemeProvider muiTheme={muiTheme}>
  <div style={styles.block}>
      <Toggle label="Simple" style={styles.toggle} />
  </div>
  </MuiThemeProvider>
  <h1>World</h1>
  </div>
  }

}

ReactDOM.render(<World/>, document.getElementById('world'));
