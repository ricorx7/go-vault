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
import CompassCalList from "./compasscalList.jsx";
import CompassCalEdit from "./compasscalEdit.jsx";
import RmaList from "./rmaList.jsx";
import RmaEdit from "./rmaEdit.jsx";
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
      <Route path="/compasscals" component={CompassCalList} />
      <Route path="/compasscals/:id" component={CompassCalEdit} />
      <Route path="/rma" component={RmaList} />
      <Route path="/rma/:id" component={RmaEdit} />
    </Route>
  </Router>
), document.getElementById('app'))