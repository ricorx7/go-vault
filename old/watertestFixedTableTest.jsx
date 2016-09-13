import React from 'react';
import ReactDOM from 'react-dom';
import {Table, Column, Cell} from 'fixed-data-table';
 
class MyTableTest extends React.Component {
 
  constructor(props) {
    super(props);
    this.rows = [{"id":1,"first_name":"William","last_name":"Elliott","email":"welliott0@wisc.edu",
             "country":"Argentina","ip_address":"247.180.226.89"},
              {"id":2,"first_name":"Carl","last_name":"Ross","email":"cross1@mlb.com",
             "country":"South Africa","ip_address":"27.146.70.36"},
              {"id":3,"first_name":"Jeremy","last_name":"Scott","email":"jscott2@cbsnews.com",
             "country":"Colombia","ip_address":"103.52.74.225"}];
    this.state = {
             filteredDataList: this.rows,
             sortBy: 'id',
             sortDir: null
    };
  }

_renderHeader(label, cellDataKey) {
  return <div>
        <a onClick={this._sortRowsBy.bind(this, cellDataKey)}>{label}</a>
          <div>
            <br />
            <input style={{width:90+'%'}} onChange={this._onFilterChange.bind(this, cellDataKey)}/>
          </div>
      </div>;
}
 
_onFilterChange(cellDataKey, event) {
  if (!event.target.value) {
    this.setState({
      filteredDataList: this.rows,
    });
  }
  var filterBy = event.target.value.toString().toLowerCase();
  var size = this.rows.length;
  var filteredList = [];
  for (var index = 0; index < size; index++) {
    var v = this.rows[index][cellDataKey];
    if (v.toString().toLowerCase().indexOf(filterBy) !== -1) {
      filteredList.push(this.rows[index]);
    }
  }
  this.setState({
    filteredDataList: filteredList,
  });
}

_sortRowsBy(cellDataKey) {
  var sortDir = this.state.sortDir;
  var sortBy = cellDataKey;
  if (sortBy === this.state.sortBy) {
    sortDir = this.state.sortDir === 'ASC' ? 'DESC' : 'ASC';
  } else {
    sortDir = 'DESC';
  }
  var rows = this.state.filteredDataList.slice();
  rows.sort((a, b) => {
    var sortVal = 0;
    if (a[sortBy] > b[sortBy]) {
      sortVal = 1;
    }
    if (a[sortBy] < b[sortBy]) {
      sortVal = -1;
    }
 
    if (sortDir === 'DESC') {
      sortVal = sortVal * -1;
    }
    return sortVal;
  });
 
  this.setState({sortBy, sortDir, filteredDataList : rows});
}

 
  render() {

var sortDirArrow = '';
if (this.state.sortDir !== null){
  sortDirArrow = this.state.sortDir === 'DESC' ? ' ↓' : ' ↑';
}

      return <Table
        height={40+((this.rows.length+1) * 30)}
        width={1150}
        rowsCount={this.rows.length}
        rowHeight={30}
        headerHeight={30}
        rowGetter={function(rowIndex) {return this.rows[rowIndex]; }.bind(this)}>
        <Column dataKey="id" width={50} label={'id' + (this.state.sortBy === 'id' ? sortDirArrow : '')} headerRenderer={this._renderHeader.bind(this)}/>
        <Column dataKey="first_name" width={200} label="First Name" />
        <Column  dataKey="last_name" width={200} label="Last Name" />
        <Column  dataKey="email" width={400} label="e-mail" />
        <Column  dataKey="country" width={300} label="Country" />
      </Table>;
  }
}
 
ReactDOM.render(<MyTableTest />, document.getElementById('myfixedtable'));