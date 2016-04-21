var AdcpRow = React.createClass({
  render: function() {
    return (
      <li key={this.props.adcp.id} className="list-group-item-info">
        <h4 class="list-group-item-heading">{this.props.adcp.SerialNumber}</h4>
         <p>{this.props.adcp.Customer}</p>
         <p>{this.props.adcp.OrderNumber}</p>
         <a href={'../adcp/update/' + this.props.adcp.SerialNumber} className="btn btn-warning">EDIT</a>
         <a href={'../adcp_checklist/update/' + this.props.adcp.SerialNumber} className="btn btn-warning">CHECKLIST</a>
         <a href={'../adcp/cert/' + this.props.adcp.SerialNumber} className="btn btn-info">CERT</a>

        <form action="../adcp/wt" method="post" role="form" className="navbar-form">
          <div class="form-group">
            <input type="hidden" className="form-control" value="{this.props.adcp.SerialNumber}" name="PartialSerialNumber" placeholder="filter SerialNumber"></input>
          </div>
          <button type="submit" className="btn btn-primary">Lake Test</button>
        </form>

        <small>{this.props.adcp.Created}</small>
      </li>
    );
  }
});

var AdcpList = React.createClass({
  render: function() {
    var rows = [];
    var filter = this.props.filterAdcp;

    this.props.adcps.forEach(function(adcp) {
      if (adcp.SerialNumber.indexOf(filter) === -1) {
        return true;  // Skip this ADCP
      }
      rows.push(<AdcpRow adcp={adcp} key={adcp.id} />);
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
      this.refs.filterAdcpInput.value,
      this.refs.isNotShippedInput.checked
    );
  },

  render: function() {
    return (
      <form>
        <input
          type="text"
          placeholder="Search ADCP..."
          value={this.props.filterAdcp}
          ref="filterAdcpInput"
          onChange={this.handleChange}
        />
        <p>
          <input
            type="checkbox"
            checked={this.props.isNotShipped}
            ref="isNotShippedInput"
            onChange={this.handleChange}
          />
          {' '}
          Only show ADCPs that have not shipped
        </p>
      </form>
    );
  }
});


var FilterableAdcpList = React.createClass({
  getInitialState: function() {
    return {
      filterAdcp: "",
      isNotShipped: false,
      data: {Adcps:[]}
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

  handleUserInput: function(filterAdcp, isNotShipped) {
      this.setState({
        filterAdcp: filterAdcp,
        isNotShipped: isNotShipped
      });
    },

  render: function() {
    return (
      <div>
        <SearchBar
          filterAdcp={this.state.filterAdcp}
          isNotShipped={this.state.isNotShipped}
          onUserInput={this.handleUserInput} />

        <AdcpList
          adcps={this.state.data.Adcps}
          filterAdcp={this.state.filterAdcp}
          isNotShipped={this.state.isNotShipped} />
      </div>
    );
  }
});

// var Adcps = React.createClass({
//   getInitialState: function() {
//     return {data: {Adcps:[]}};
//   },
//   componentDidMount: function() {
//     $.ajax({
//       url: this.props.url,
//       dataType: 'json',
//       cache: false,
//       success: function(data) {
//         this.setState({data: data});
//       }.bind(this),
//       error: function(xhr, status, err) {
//         console.error(this.props.url, status, err.toString());
//       }.bind(this)
//     });
//   },
//   render: function() {
//     var adcpList = this.state.data.Adcps.map(function (adcp) {
//          return (
//              <li key={adcp.id} className="list-group-item">
//                <h4 class="list-group-item-heading">{adcp.SerialNumber}</h4>
//                 <p>{adcp.Customer}</p>
//                 <p>{adcp.OrderNumber}</p>
//                 <a href={'../adcp/update/' + adcp.SerialNumber} className="btn btn-warning">EDIT</a>
//                 <a href={'../adcp_checklist/update/' + adcp.SerialNumber} className="btn btn-warning">CHECKLIST</a>
//                 <a href={'../adcp/cert/' + adcp.SerialNumber} className="btn btn-info">CERT</a>
//
//                <form action="../adcp/wt" method="post" role="form" className="navbar-form">
//                  <div class="form-group">
//                    <input type="hidden" className="form-control" value="{adcp.SerialNumber}" name="PartialSerialNumber" placeholder="filter SerialNumber"></input>
//                  </div>
//                  <button type="submit" className="btn btn-primary">Lake Test</button>
//                </form>
//
//                <small>{adcp.Created}</small>
//              </li>
//          );
//        });
//
//     return <div>{adcpList}</div>;
//   }
// });

// var Adcps = React.createClass({
//   getInitialState: function() {
//     return {data: {Adcps:[]}};
//   },
//   componentDidMount: function() {
//     $.ajax({
//       url: this.props.url,
//       dataType: 'json',
//       cache: false,
//       success: function(data) {
//         this.setState({data: data});
//       }.bind(this),
//       error: function(xhr, status, err) {
//         console.error(this.props.url, status, err.toString());
//       }.bind(this)
//     });
//   },
//   render: function() {
//     var adcpList = this.state.data.Adcps.map(function (adcp) {
//          return (
//              <li key={adcp.id} className="list-group-item">
//                <h4 class="list-group-item-heading">{adcp.SerialNumber}</h4>
//                 <p>{adcp.Customer}</p>
//                 <p>{adcp.OrderNumber}</p>
//                 <a href={'../adcp/update/' + adcp.SerialNumber} className="btn btn-warning">EDIT</a>
//                 <a href={'../adcp_checklist/update/' + adcp.SerialNumber} className="btn btn-warning">CHECKLIST</a>
//                 <a href={'../adcp/cert/' + adcp.SerialNumber} className="btn btn-info">CERT</a>
//
//                <form action="../adcp/wt" method="post" role="form" className="navbar-form">
//                  <div class="form-group">
//                    <input type="hidden" className="form-control" value="{adcp.SerialNumber}" name="PartialSerialNumber" placeholder="filter SerialNumber"></input>
//                  </div>
//                  <button type="submit" className="btn btn-primary">Lake Test</button>
//                </form>
//
//                <small>{adcp.Created}</small>
//              </li>
//          );
//        });
//
//     return <div>{adcpList}</div>;
//   }
// });


ReactDOM.render(
  <FilterableAdcpList url="/vault/adcp" />,
  document.getElementById('adcpList')
);
