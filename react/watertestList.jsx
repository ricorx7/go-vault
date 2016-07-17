import React from 'react';
import ReactDOM from 'react-dom';
import Bootstrap from 'bootstrap'
import {blueGrey500} from 'material-ui/styles/colors';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import Toggle from 'material-ui/toggle';
import FontIcon from 'material-ui/FontIcon';
import FlatButton from 'material-ui/FlatButton';
import RaisedButton from 'material-ui/RaisedButton';
import Checkbox from 'material-ui/Checkbox';
import ReactPaginate from 'react-paginate';

const muiTheme = getMuiTheme({
  palette: {
    accent1Color: blueGrey500,
  },
});

const style = {
  margin: 12,
};

var WaterTestRow = React.createClass({

  getInitialState() {
    return {
      checked: false
    };
  },

  handleChange(checked) {
  this.setState({
    checked: checked
  });
},

  handleIsSelectedChange: function(wt) {
    // Use API to change the value
    console.log(wt.id)
  },

  render: function() {
    var isSelected;
    if(this.props.wt.IsSelected) { isSelected = "Selected" }
    else { isSelected = ""}

    const styles = {
      block: {
        maxWidth: 150,
      },
      toggle: {
        marginBottom: 10,
      },
    };

    return (
      <div>
      <li key={this.props.wt.id}  className="list-group-item">
        <div className="row">
          <div className="col-xs-5">
            <span>{this.props.wt.SerialNumber}</span>
          </div>
          <div className="col-xs-4">
            <span>{this.props.wt.SubsystemDescStr}</span>
          </div>
        </div>

        <div className="row">
          <div className="col-xs-5">
            <strong> Beam <span>{this.props.wt.TestOrientation}</span> Forward</strong>
          </div>
          <div className="col-xs-5">
            <strong> {isSelected} </strong>
            <MuiThemeProvider muiTheme={muiTheme}>
              <div style={styles.block}>
                  <Toggle label="Is Selected" defaultToggled={this.props.wt.IsSelected} onToggle={this.handleIsSelectedChange.bind(this, this.props.wt)}  style={styles.toggle} />
              </div>
            </MuiThemeProvider>
          </div>
        </div>

        <div className="row">
          <table className="table table-striped table-condensed">
            <tbody>
              <tr>
                <td><strong>GPS</strong></td>
                <td className="text-center"><span>{this.props.wt.GpsDistance}</span></td>
                <td className="text-center"><span>{this.props.wt.GpsDirection}</span></td>
                <td className="text-center"><span></span></td>
                <td className="text-center"><span></span></td>
              </tr>
              <tr>
                <td><strong>DMG</strong></td>
                <td className="text-center"><span>{this.props.wt.BtDistance}</span></td>
                <td className="text-center"><span>{this.props.wt.BtDirection}</span></td>
                <td className="text-center"><span>{this.props.wt.DistanceError}</span></td>
                <td className="text-center"><span>{this.props.wt.DirectionError}</span></td>
              </tr>
            </tbody>
          </table>
        </div>

        <div className="row">
          <div className="col-xs-5">
            <MuiThemeProvider muiTheme={muiTheme}>
              <FlatButton label="Edit" linkButton={true} primary={true} target="_blank" href={'wt/update/'+ this.props.wt.id} />
            </MuiThemeProvider>
            <MuiThemeProvider muiTheme={muiTheme}>
              <FlatButton label="Report" linkButton={true} secondary={true} target="_blank" href={this.props.wt.PlotReport} />
            </MuiThemeProvider>
          </div>
        </div>

        <div className="row">
          <div className="col-xs-5">
            <small >
              Posted on
              <span>{this.props.wt.Created}</span>
              by
              <span>{this.props.wt.UserName}</span>
            </small>
          </div>
        </div>
      </li>
      </div>
    );
  }
});

var WaterTestList = React.createClass({

  render: function() {
    var rows = [];
    var filter = this.props.filter;
    var isSelected = this.props.isSelected

    var wtData = this.props.watertests;

    if(filter.length > 0) {
      wtData = this.props.watertests.filter(function(l){
                return l.SerialNumber.match( filter );
            });
      }

      wtData.forEach(function(wt) {
     //this.props.watertests.forEach(function(wt) {
    //   if (wt.SerialNumber.indexOf(filter) === -1) {
    //     return true;  // Skip this ADCP
    //   }

      //if(wt.IsSelected ==! isSelected) {
      //  return true;
      //}

      rows.push(<WaterTestRow wt={wt} key={wt.id} />);
    });
    return (
      <ul className="list-group">
        {rows}
      </ul>
    );
  }
});

var SearchBar = React.createClass({

  getInitialState: function() {
    return {searchInput: ''};
  },

  handleCheckedChange: function(event, checked) {
    this.props.onCheckSelectedInput( checked );
  },

  handleSearchChange: function(event) {
    //this.props.onSearchInput( this.refs.filterWtInput.value );
    this.props.onSearchInput( this.state.searchInput );
  },

  handleSearchInput: function(event) {
    this.setState({searchInput: event.target.value})

    // if(event.key === 'Enter') {
    //   console.log("enter pressed");
    //   this.props.onSearchInput( this.state.searchInput );
    // }
  },

  render: function() {

    const styles = {
      block: {
        maxWidth: 150,
      },
      toggle: {
        marginBottom: 10,
      },
    };

    return (
      <form>
        <input
          type="text"
          placeholder="Search ADCP..."
          value={this.state.searchInput}
          onChange={this.handleSearchInput}
        />

        <MuiThemeProvider muiTheme={muiTheme}>
          <RaisedButton label="Filter" primary={true} style={style} onClick={this.handleSearchChange} />
        </MuiThemeProvider>

      </form>
    );
  }
});


var FilterableWaterTestList = React.createClass({
  getInitialState: function() {
    return {
      filter: "",
      isSelected: true,
      offset: 0,
      data: {WaterTests:[]},
    };
  },

  componentDidMount: function() {
    this.loadWaterTestFromServer();
  },

  loadWaterTestFromServer: function() {
    $.ajax({
      url: this.props.url,
      data     : {limit: this.props.perPage, offset: this.state.offset, filter: this.state.filter},
      dataType: 'json',
      cache: false,
      success: function(data) {
        this.setState({data: data});
      }.bind(this),
      error: function(xhr, status, err) {
        console.error(this.props.url, status, err.toString());
      }.bind(this)
    });
  },

  loadWaterTestFromServerFilter: function() {
    $.ajax({
      url: this.props.url,
      data     : {filter: this.state.filter},
      dataType: 'json',
      cache: false,
      success: function(data) {
        this.setState({data: data});
      }.bind(this),
      error: function(xhr, status, err) {
        console.error(this.props.url, status, err.toString());
      }.bind(this)
    });
    console.log("filtered data")
  },

  handleUserInput: function(isSelected) {
      this.setState({
        isSelected: isSelected
      });
    },

  handleSearchInput: function(filterString) {
      this.setState({ filter: filterString});

      if(filterString == '')
      {
        this.loadWaterTestFromServer();
        console.log("Empty filter")
      }
      else
      {
        this.loadWaterTestFromServerFilter();
      }

      // var filteredData = this.state.data.WaterTests.filter(function(l){
      //           return l.SerialNumber.match( filterString );
      //       });
    },

  handlePageClick: function(data) {
    var selected = data.selected;
    var offset = Math.ceil(selected * this.props.perPage);

    this.setState({offset: offset}, () => {
      this.loadWaterTestFromServer();
    });
  },

  render: function() {
    return (
      <div>
        <SearchBar
          filter={this.state.filter}
          isSelected={this.state.isSelected}
          onSearchInput={this.handleSearchInput}
          onCheckSelectedInput={this.handleUserInput} />

        <WaterTestList
          watertests={this.state.data.WaterTests}
          filter={this.state.filter}
          isSelected={this.state.isSelected} />

          <ReactPaginate previousLabel={"previous"}
                           nextLabel={"next"}
                           breakLabel={<a href="">...</a>}
                           pageNum={this.state.pageNum}
                           marginPagesDisplayed={2}
                           pageRangeDisplayed={5}
                           clickCallback={this.handlePageClick}
                           containerClassName={"pagination"}
                           subContainerClassName={"pages pagination"}
                           activeClassName={"active"} />
      </div>
    );
  }
});


ReactDOM.render(
  <FilterableWaterTestList url="/vault/wt" perPage={10} />,
  document.getElementById('watertestList')
);
