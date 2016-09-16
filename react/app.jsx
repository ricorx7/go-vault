import React from 'react';
import ReactDOM from 'react-dom';
import { Router, Route, Link, hashHistory } from 'react-router';

import Navigation from "./header.jsx";
import AdcpList from "./adcpList.jsx";
import AdcpCert from "./adcpCert.jsx";
import WaterTestList from "./watertestList.jsx";
import WaterTestEdit from "./watertestEdit.jsx";
import TankTestList from "./tanktestList.jsx";
import TankTestEdit from "./tanktestEdit.jsx";
import SnrTestList from "./snrtestList.jsx";
import SnrTestEdit from "./snrtestEdit.jsx";
import Layout from "./Layout.jsx";

const app = document.getElementById('app');

// Finally, we render a <Router> with some <Route>s.
// It does all the fancy routing stuff for us.
ReactDOM.render((
  <Router history={hashHistory}>
    <Route path="/" component={Layout}>
      <Route path="/adcps" component={AdcpList} />
      <Route path="/adcps/cert/:id" component={AdcpCert} />
      <Route path="/watertests" component={WaterTestList} />
      <Route path="/watertests/:id" component={WaterTestEdit} />
      <Route path="/tanktests" component={TankTestList} />
      <Route path="/tanktests/:id" component={TankTestEdit} />
      <Route path="/snrtests" component={SnrTestList} />
      <Route path="/snrtests/:id" component={SnrTestEdit} />
    </Route>
  </Router>
), document.getElementById('app'))