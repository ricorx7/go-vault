import React from 'react';
import ReactDOM from 'react-dom';


var Header = React.createClass({
  render: function() {
    return (
      <div>
        <nav className="navbar navbar-inverse">
          <div className="container-fluid">

            <div className="navbar-header">
              <button type="button" className="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1" aria-expanded="false">
                <span className="sr-only">Toggle navigation</span>
                <span className="icon-bar"></span>
                <span className="icon-bar"></span>
                <span className="icon-bar"></span>
              </button>
              <a className="navbar-brand" href="/adcp">RoweTech Inc. Vault</a>
            </div>

            <div className="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
              <ul className="nav navbar-nav">
                <li className="active"><a href="/adcp">ADCP <span class="sr-only"></span></a></li>
                <li className="dropdown">
                  <a href="/adcp" className="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">ADCP<span class="caret"></span></a>
                  <ul className="dropdown-menu">
                    <li><a href="/adcp">List</a></li>
                    <li role="separator" class="divider"></li>
                    <li><a href="/adcp/add">Add</a></li>
                  </ul>
                </li>
                <li className="dropdown">
                  <a href="/adcp/wt" className="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Lake Tests<span class="caret"></span></a>
                  <ul className="dropdown-menu">
                    <li><a href="/adcp/wt">List</a></li>
                    <li role="separator" class="divider"></li>
                    <li><a href="/adcp/wt/add">Add</a></li>
                  </ul>
                </li>
                <li className="dropdown">
                  <a href="/rma" className="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">RMA<span class="caret"></span></a>
                  <ul className="dropdown-menu">
                    <li><a href="/rma">List</a></li>
                    <li role="separator" class="divider"></li>
                    <li><a href="/rma/add">Add</a></li>
                  </ul>
                </li>
                <li className="dropdown">
                  <a href="/so" className="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Sales Order<span class="caret"></span></a>
                  <ul className="dropdown-menu">
                    <li><a href="/so">List</a></li>
                    <li role="separator" class="divider"></li>
                    <li><a href="/so/add">Add</a></li>
                  </ul>
                </li>
                <li className="dropdown">
                  <a href="/product" className="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Product<span class="caret"></span></a>
                  <ul className="dropdown-menu">
                    <li><a href="/product">List</a></li>
                    <li role="separator" class="divider"></li>
                    <li><a href="/product/add">Add</a></li>
                  </ul>
                </li>
              </ul>
              <ul className="nav navbar-nav navbar-right">
                <li><a target="_blank" href="http://rowetechinc.co/wiki/index.php?title=Main_Page" Redirect>RoweTech Wiki</a></li>
              </ul>
            </div>
          </div>
        </nav>
        </div>

    );
  }
});

ReactDOM.render(
  <Header />,
  document.getElementById('header')
);
