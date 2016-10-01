import React from 'react';
import ReactDOM from 'react-dom';
import {DataTable} from 'react-data-components';
import {blueGrey500} from 'material-ui/styles/colors';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import Toggle from 'material-ui/toggle';
import { Button, Row, Col, Table, Glyphicon, Form, FormControl, FormGroup, ControlLabel, HelpBlock, Well, Checkbox, Panel  } from 'react-bootstrap';
import DatePicker from 'react-datepicker';

// Theme for material-ui toggle
const muiTheme = getMuiTheme({
  palette: {
    accent1Color: blueGrey500,
  },
});

// Style for material-ui toggle
const styles = {
  block: {
    maxWidth: 150,
  },
  toggle: {
    marginBottom: 10,
  },
};


// Edit the Tank test data.
export default class RmaEdit extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
        data: {},
    }
  }

    // At startup get all the Tank Test data
  componentDidMount() {
    this.apiGetRmaSelected();
    console.log("data length %i\n", this.state.data);
  }

    // Call API to set IsSelect selection
    apiGetRmaSelected() {
      var urlSelected = "/vault/rma/edit/" + this.props.params.id;
      $.ajax({
        url: urlSelected,
        dataType: 'json',
        cache: false,
        success: function(data) {
          console.log("Data gotten from %s\n", urlSelected);
          console.log(data);
          this.setState({data: data});
        }.bind(this),
        error: function(xhr, status, err) {
          console.error(urlSelected, status, err.toString());
        }.bind(this)
      });
    }

    // Call API to set IsSelect selection
    apiPostRma(wtData) {
      var urlSelected = "/vault/rma/edit/" + this.state.data.id;
      $.ajax({
        url: urlSelected,
        dataType: 'json',
        type: 'POST',
        data: wtData,
        cache: false,
        success: function(data) {
          console.log("Data updated from %s\n", urlSelected);
        }.bind(this),
        error: function(xhr, status, err) {
          console.error(urlSelected, status, err.toString());
        }.bind(this)
      });
    }

    // Update the DB with the latest data
    updateDB()
    {
        console.log(this.state.data);
        this.apiPostRma(JSON.stringify(this.state.data));
    }

    getValidationState() {
        const length = this.state.gpsDistance;
        if (length > 10) return 'success';
        else if (length > 5) return 'warning';
        else if (length > 0) return 'error';
    }

    // Update the RMA number
	rmaNumberChange(e) {
        this.state.data.RmaNumber = e.target.value;             // Update the object
        this.update();                                            // Update DB and display 
    }

    // Set the RMA Type.
    rmaTypeChange(e) {
        this.state.data.RmaType = e.target.value;               // Update the object
        this.update();                                              // Update DB and display   
    }

    // Set the Customer.
    companyChange(e) {
        this.state.data.Company = e.target.value;               // Update the object
        this.update();                                              // Update DB and display   
    }

    // Set the Customer.
    rmaDateChange(e) {
        this.state.data.RmaDate = e.target.value;               // Update the object
        this.update();                                              // Update DB and display   
    }

    // Set the test orientation.
    notesChange(e) {
        this.state.data.Notes = e.target.value;                    // Update the object
        this.update();                                             // Update DB and display   
    }

    // Update the DB and display.
    update() {
        this.forceUpdate();                                       // Refresh display
        this.updateDB();                                          // Update the database
    }

    // Convert to Bool
    convertToBool(val) {
      return (val === true);
    } 


  // Render function
  render() {
    // Waiting for AJAX response
    if(this.state.data == null)
    {
        return(<div>Loading...</div>);
    }

    const containerStyle = {
      margin: "10px"
    };

    return (
        <div style={containerStyle}>
          <Well><h1>{this.state.data.RmaType}{this.state.data.RmaNumber} - {this.state.data.Company}</h1></Well>
          
          <Row>
          <Panel header="RMA Info" bsStyle="info" collapsible defaultExpanded >
            <Form horizontal>
              <FormGroup controlId="formControlsSelect">
                <Col sm={2}>
                    <ControlLabel>RMA Type:</ControlLabel>
                </Col>
                <Col sm={10}>
                    <FormControl componentClass="select" placeholder="RmaType" value={this.state.data.RmaType} onChange={this.rmaTypeChange.bind(this)}>
                    <option value="290">290 - Warranty</option>
                    <option value="280">280 - Billable</option>
                    <option value="259">259 - Demo Repair</option>
                    </FormControl>
                </Col>
              </FormGroup>

              <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                <Col sm={2}>
                    <ControlLabel>RMA Number:</ControlLabel>
                </Col>
                <Col sm={10}>
                    <FormControl type="text" value={this.state.data.RmaNumber} placeholder="Enter text" onChange={this.rmaNumberChange.bind(this)} />
                    <FormControl.Feedback /> 
                </Col>
              </FormGroup>

              <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                <Col sm={2}>
                    <ControlLabel>Company:</ControlLabel>
                </Col>
                <Col sm={10}>
                    <FormControl type="text" value={this.state.data.Company} placeholder="Enter text" onChange={this.companyChange.bind(this)} />
                    <FormControl.Feedback />
                </Col> 
              </FormGroup>

              <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                <Col sm={2}>
                    <ControlLabel>Date:</ControlLabel>
                </Col>
                <Col sm={10}>
                    <FormControl type="text" value={this.state.data.RmaDate} placeholder="Enter text" onChange={this.rmaDateChange.bind(this)} />
                    <FormControl.Feedback />
                    <DatePicker selected={this.state.data.RmaDate}  onChange={this.rmaDateChange.bind(this)} />
                </Col> 
              </FormGroup>

            </Form>
        </Panel>
        </Row>

          <Row >
            <Col xs={8}>
              <FormGroup controlId="formControlsTextarea">
                <ControlLabel>Notes:</ControlLabel>
                <FormControl componentClass="textarea" placeholder="Notes" value={this.state.data.Notes} onChange={this.notesChange.bind(this)} />
              </FormGroup>
            </Col>
          </Row>
        
        </div>
    );
  }
}

// Use the url PROP to get the Water Test data
//ReactDOM.render((<WaterTestEdit url="/vault/wt/edit/" />), document.getElementById('watertestEdit'));
