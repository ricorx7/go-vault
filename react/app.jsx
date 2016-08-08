import React from 'react';
import ReactDOM from 'react-dom';
import { Router, Route, Link, hashHistory } from 'react-router';

import Navigation from "./header.jsx";
import AdcpList from "./adcpList.jsx"
import WaterTestCompList from "./watertestListComp.jsx";
import WaterTestEdit from "./watertestEdit.jsx";
import Layout from "./Layout.jsx";

const app = document.getElementById('app');

// Finally, we render a <Router> with some <Route>s.
// It does all the fancy routing stuff for us.
ReactDOM.render((
  <Router history={hashHistory}>
    <Route path="/" component={Layout}>
      <Route path="/adcps" component={AdcpList} />
      <Route path="/adcps/cert/:id" component={AdcpList} />
      <Route path="/watertests" component={WaterTestCompList} />
      <Route path="/watertests/:id" component={WaterTestEdit} />
    </Route>
  </Router>
), document.getElementById('app'))