var WaterTestRow = React.createClass({
  render: function() {
    var isSelected;
    if(this.props.wt.IsSelected) { isSelected = "SELECTED" }
    else { isSelected = ""}

    return (
      <li key={this.props.wt.id}  className="list-group-item">
        <div className="row">
          <div className="col-xs-5">
            <span>{this.props.wt.SerialNumber}</span>
          </div>
          <div className="col-xs-4">
            <span>{this.props.wt.SubsystemDescStr}</span>
          </div>
          <div className="col-xs-3">
            <a href={'wt/update/'+ this.props.wt.id} className="btn btn-info">Edit</a>
          </div>
        </div>

        <div className="row">
          <div className="col-xs-5">
            <strong> Beam <span>{this.props.wt.TestOrientation}</span> Forward</strong>
          </div>
          <div className="col-xs-5">
            <strong> {isSelected} </strong>
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
            <a href={this.props.wt.PlotReport} target="_blank" className="btn btn-info">Report</a>
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
    );
  }
});

var WaterTestList = React.createClass({
  render: function() {
    var rows = [];
    var filter = this.props.filter;
    var isSelected = this.props.isSelected

    this.props.watertests.forEach(function(wt) {
      if (wt.SerialNumber.indexOf(filter) === -1) {
        return true;  // Skip this ADCP
      }

      if(wt.IsSelected ==! isSelected) {
        return true;
      }

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

  handleChange: function() {
    this.props.onUserInput(
      this.refs.filterWtInput.value,
      this.refs.isSelectedInput.checked
    );
  },

  render: function() {
    return (
      <form>
        <input
          type="text"
          placeholder="Search ADCP..."
          value={this.props.filter}
          ref="filterWtInput"
          onChange={this.handleChange}
        />
        <p>
          <input
            type="checkbox"
            checked={this.props.isSelected}
            ref="isSelectedInput"
            onChange={this.handleChange}
          />
          {' '}
          Only show WaterTests that are selected
        </p>
      </form>
    );
  }
});


var FilterableWaterTestList = React.createClass({
  getInitialState: function() {
    return {
      filter: "",
      isSelected: true,
      data: {WaterTests:[]}
    };
  },

  componentDidMount: function() {
    $.ajax({
      url: this.props.url,
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

  handleUserInput: function(filter, isSelected) {
      this.setState({
        filter: filter,
        isSelected: isSelected
      });
    },

  render: function() {
    return (
      <div>
        <SearchBar
          filter={this.state.filter}
          isSelected={this.state.isSelected}
          onUserInput={this.handleUserInput} />

        <WaterTestList
          watertests={this.state.data.WaterTests}
          filter={this.state.filter}
          isSelected={this.state.isSelected} />
      </div>
    );
  }
});


ReactDOM.render(
  <FilterableWaterTestList url="/vault/wt" />,
  document.getElementById('watertestList')
);
