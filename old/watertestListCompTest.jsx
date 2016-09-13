const React = require('react');
const ReactDOM = require('react-dom');
var DataTable = require('react-data-components').DataTable;

// Generate random data
var names = [ 'Carlos', 'Juan', 'Jesus', 'Alberto', 'John' ];
var cities = [ 'Chicago', 'Tampico', 'San Francisco', 'Mexico City', 'Boston', 'New York' ];
var addresses = [ '333 West Wacker Drive', '1931 Insurgentes Sur', '1 Lombard Street', '55 Av Hidalgo'];

var data = [];
for (var i = 0; i < 1000; i++) {
  data.push({
    id: i,
    name: names[~~(Math.random() * names.length)],
    city: cities[~~(Math.random() * cities.length)],
    address: addresses[~~(Math.random() * addresses.length)]
  });
}

var columns = [
  { title: 'Name', prop: 'name'  },
  { title: 'City', prop: 'city' },
  { title: 'Address', prop: 'address' }
];

ReactDOM.render((
    <DataTable
      className="container"
      keys="id"
      columns={columns}
      initialData={data}
      initialPageLength={5}
      initialSortBy={{ prop: 'city', order: 'descending' }}
      pageLengthOptions={[ 5, 20, 50 ]}
    />
  ), document.getElementById('compTableTest'));

