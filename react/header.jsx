import React from 'react';
import ReactDOM from 'react-dom';
import { Nav, Navbar, NavbarCollapse, NavbarBrand, NavbarHeader, NavbarToggle, NavDropdown, NavItem, MenuItem, FormGroup, FormControl, Button } from 'react-bootstrap';

export default class Navigation extends React.Component {

  constructor(props) {
    super(props);

    this.state = {
        index: 0
    };
  }

handleSelect(eventKey) {
    console.log(eventKey);
    switch(eventKey) {
        case 1.1:
            console.log("ADCP List");
            break;
        case 1.2:
            console.log("ADCP Add");
            break;
        case 2.1:
            console.log("Water Test Add");
            break;
        case 2.2:
            console.log("Water Test Add");
            break;
        case 3.1:
            console.log("Tank Test Add");
            break;
        case 3.2:
            console.log("Tank Test Add");
            break;
        default:
            break;
    }
}

 render() {
    return (
    <Navbar inverse>
        <Navbar.Header>
            <Navbar.Brand>
                <a href="#">RoweTech Inc.Vault</a>
            </Navbar.Brand>
            <Navbar.Toggle />
        </Navbar.Header>
        <Navbar.Collapse>
            <Nav onSelect={this.handleSelect}>
                <NavDropdown eventKey={1} title="ADCP" id="basic-nav-dropdown">
                    <MenuItem eventKey={1.1} href="">List</MenuItem>
                    <MenuItem eventKey={1.2}>Add</MenuItem>
                </NavDropdown>
                <NavDropdown eventKey={2} title="WaterTest" id="basic-nav-dropdown">
                    <MenuItem eventKey={2.1} href="watertestListComp.html">List</MenuItem>
                    <MenuItem eventKey={2.2}>Add</MenuItem>
                </NavDropdown>
                <NavDropdown eventKey={3} title="TankTest" id="basic-nav-dropdown">
                    <MenuItem eventKey={3.1} href="tanktestList.html">List</MenuItem>
                    <MenuItem eventKey={3.2}>Add</MenuItem>
                </NavDropdown>
                <NavDropdown eventKey={4} title="SNR Test" id="basic-nav-dropdown">
                    <MenuItem eventKey={4.1} href="">List</MenuItem>
                    <MenuItem eventKey={4.2}>Add</MenuItem>
                </NavDropdown>
                <NavDropdown eventKey={5} title="RMA" id="basic-nav-dropdown">
                    <MenuItem eventKey={5.1} href="">List</MenuItem>
                    <MenuItem eventKey={5.2}>Add</MenuItem>
                </NavDropdown>
                <NavDropdown eventKey={6} title="Sales Order" id="basic-nav-dropdown">
                    <MenuItem eventKey={6.1} href="">List</MenuItem>
                    <MenuItem eventKey={6.2}>Add</MenuItem>
                </NavDropdown>
                <NavDropdown eventKey={7} title="Product" id="basic-nav-dropdown">
                    <MenuItem eventKey={7.1} href="">List</MenuItem>
                    <MenuItem eventKey={7.2}>Add</MenuItem>
                </NavDropdown>
            </Nav>
            <Nav pullRight>
                <NavItem eventKey={1} href="http://rowetechinc.co/wiki/index.php?title=Main_Page">RoweTech Wiki</NavItem>
            </Nav>
        </Navbar.Collapse>
    </Navbar>
    )}

};

//ReactDOM.render(navbarInstance, mountNode);
//ReactDOM.render(<Navigation />, document.getElementById('header'));