import React from 'react';
import ReactDOM from 'react-dom';
import {DataTable} from 'react-data-components';
import {blueGrey500} from 'material-ui/styles/colors';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import Toggle from 'material-ui/toggle';
import { Button, Row, Col, Table, Glyphicon, Form, FormControl, FormGroup, ControlLabel, HelpBlock, Well, Checkbox, Panel  } from 'react-bootstrap';

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


// Edit the Compass Cal data.
export default class CompassCalEdit extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
        data: {},
    }
  }

    // At startup get all the Compass Cal data
  componentDidMount() {
    this.apiGetCompassCalSelected();
    console.log("data length %i\n", this.state.data);
  }

    // Call API to set IsSelect selection
    apiGetCompassCalSelected() {
      var urlSelected = "/vault/compasscal/edit/" + this.props.params.id;
      $.ajax({
        url: urlSelected,
        dataType: 'json',
        cache: false,
        success: function(data) {
          console.log("Data gotten from %s\n", urlSelected);
          console.log(data);
          this.setState({data: data});
          //this.init();
        }.bind(this),
        error: function(xhr, status, err) {
          console.error(urlSelected, status, err.toString());
        }.bind(this)
      });
    }

    // Call API to set IsSelect selection
    apiPostCompassCal(ccData) {
      var urlSelected = "/vault/compasscal/edit/" + this.state.data.id;
      $.ajax({
        url: urlSelected,
        dataType: 'json',
        type: 'POST',
        data: ccData,
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
        this.apiPostCompassCal(JSON.stringify(this.state.data));
    }

    getValidationState() {
        const length = this.state.gpsDistance;
        if (length > 10) return 'success';
        else if (length > 5) return 'warning';
        else if (length > 0) return 'error';
    }

    // Update the serial number
	  serialNumberChange(e) {
        this.state.data.SerialNumber = e.target.value;             // Update the object
        this.update();                                            // Update DB and display 
    }

    // Set the test orientation.
    orientationChange(e) {
        this.state.data.TestOrientation = parseInt(e.target.value); // Update the object
        this.update();                                              // Update DB and display   
    }

    // Set the test orientation.
    notesChange(e) {
        this.state.data.Notes = e.target.value;                    // Update the object
        this.update();                                             // Update DB and display   
    }

    point1PreHdgChange(e) {
        this.state.data.Point1_Pre_Hdg = parseFloat(e.target.value);  // Update the object
        this.update();                                             // Update DB and display   
    }

    point2PreHdgChange(e) {
        this.state.data.Point2_Pre_Hdg = parseFloat(e.target.value);  // Update the object
        this.update();                                             // Update DB and display   
    }

    point3PreHdgChange(e) {
        this.state.data.Point3_Pre_Hdg = parseFloat(e.target.value);  // Update the object
        this.update();                                             // Update DB and display   
    }

    point4PreHdgChange(e) {
        this.state.data.Point4_Pre_Hdg = parseFloat(e.target.value);  // Update the object
        this.update();                                             // Update DB and display   
    }

    point1PrePtchChange(e) {
        this.state.data.Point1_Pre_Ptch = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    point2PrePtchChange(e) {
        this.state.data.Point2_Pre_Ptch = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    point3PrePtchChange(e) {
        this.state.data.Point3_Pre_Ptch = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    point4PrePtchChange(e) {
        this.state.data.Point4_Pre_Ptch = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    point1PreRollChange(e) {
        this.state.data.Point1_Pre_Roll = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    point2PreRollChange(e) {
        this.state.data.Point2_Pre_Roll = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    point3PreRollChange(e) {
        this.state.data.Point3_Pre_Roll = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    point4PreRollChange(e) {
        this.state.data.Point4_Pre_Roll = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    point1PostHdgChange(e) {
        this.state.data.Point1_Post_Hdg = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    point2PostHdgChange(e) {
        this.state.data.Point2_Post_Hdg = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    point3PostHdgChange(e) {
        this.state.data.Point3_Post_Hdg = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    point4PostHdgChange(e) {
        this.state.data.Point4_Post_Hdg = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    point1PostPtchChange(e) {
        this.state.data.Point1_Post_Ptch = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    point2PostPtchChange(e) {
        this.state.data.Point2_Post_Ptch = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    point3PostPtchChange(e) {
        this.state.data.Point3_Post_Ptch = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    point4PostPtchChange(e) {
        this.state.data.Point4_Post_Ptch = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    point1PostRollChange(e) {
        this.state.data.Point1_Post_Roll = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    point2PostRollChange(e) {
        this.state.data.Point2_Post_Roll = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    point3PostRollChange(e) {
        this.state.data.Point3_Post_Roll = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    point4PostRollChange(e) {
        this.state.data.Point4_Post_Roll = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    // Set IsSelected.
    isSelectedChange(e) {
        if(e === true) {
           this.state.data.IsSelected = false;                    // Invert 
         } else {
           this.state.data.IsSelected = true;
         } 

        this.update();                                            // Update DB and display
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

    return (
        <div>
          <Well><h1>{this.state.data.SerialNumber}</h1></Well>
          
          <Row>
            <Col xs={5}>
              <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                <ControlLabel>SerialNumber:</ControlLabel>
                <FormControl type="text" value={this.state.data.SerialNumber} placeholder="Enter text" onChange={this.serialNumberChange.bind(this)} />
                <FormControl.Feedback /> 
              </FormGroup>
            </Col>
          </Row>

          <Row>
            <Col xs={3}>
              <MuiThemeProvider muiTheme={muiTheme}>
                <div style={styles.block}>
                    <Toggle label="Is Selected:" defaultToggled={this.convertToBool(this.state.data.IsSelected)} onToggle={this.isSelectedChange.bind(this, this.state.data.IsSelected)} style={styles.toggle} />
                </div>
              </MuiThemeProvider>
            </Col>
          </Row>

          <Row>
            <Col xs={10}>
              <Table striped bordered condensed responsive hover>
                <thead>
                  <tr>
                    <th colSpan="5">Post Points</th>
                  </tr>
                  <tr>
                    <th></th>
                    <th>Beam 0</th>
                    <th>Beam 1</th>
                    <th>Beam 2</th>
                    <th>Beam 3</th>
                  </tr>
                </thead>
                <tbody>
                  <tr>
                    <td>Hdg</td>
                    <td>                  
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Point1_Post_Hdg} placeholder="Enter text" onChange={this.point1PostHdgChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Point2_Post_Hdg} placeholder="Enter text" onChange={this.point2PostHdgChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Point3_Post_Hdg} placeholder="Enter text" onChange={this.point3PostHdgChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Point4_Post_Hdg} placeholder="Enter text" onChange={this.point4PostHdgChange.bind(this)} />
                      </FormGroup>
                    </td>
                  </tr>
                  <tr>
                    <td>Pitch</td>
                    <td>                  
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Point1_Post_Ptch} placeholder="Enter text" onChange={this.point1PostPtchChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Point2_Post_Ptch} placeholder="Enter text" onChange={this.point2PostPtchChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Point3_Post_Ptch} placeholder="Enter text" onChange={this.point3PostPtchChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Point4_Post_Ptch} placeholder="Enter text" onChange={this.point4PostPtchChange.bind(this)} />
                      </FormGroup>
                    </td>
                  </tr>
                  <tr>
                    <td>Roll</td>
                    <td>                  
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Point1_Post_Roll} placeholder="Enter text" onChange={this.point1PostRollChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Point2_Post_Roll} placeholder="Enter text" onChange={this.point2PostRollChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Point3_Post_Roll} placeholder="Enter text" onChange={this.point3PostRollChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Point4_Post_Roll} placeholder="Enter text" onChange={this.point4PostRollChange.bind(this)} />
                      </FormGroup>
                    </td>
                  </tr>
                </tbody>
              </Table>
            </Col>
          </Row>

          <Row>
            <Col xs={10}>
              <Table striped bordered condensed responsive hover>
                <thead>
                  <tr>
                    <th colSpan="5">Pre Points</th>
                  </tr>
                  <tr>
                    <th></th>
                    <th>Beam 0</th>
                    <th>Beam 1</th>
                    <th>Beam 2</th>
                    <th>Beam 3</th>
                  </tr>
                </thead>
                <tbody>
                  <tr>
                    <td>Hdg</td>
                    <td>                  
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Point1_Pre_Hdg} placeholder="Enter text" onChange={this.point1PreHdgChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Point2_Pre_Hdg} placeholder="Enter text" onChange={this.point2PreHdgChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Point3_Pre_Hdg} placeholder="Enter text" onChange={this.point3PreHdgChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Point4_Pre_Hdg} placeholder="Enter text" onChange={this.point4PreHdgChange.bind(this)} />
                      </FormGroup>
                    </td>
                  </tr>
                  <tr>
                    <td>Pitch</td>
                    <td>                  
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Point1_Pre_Ptch} placeholder="Enter text" onChange={this.point1PrePtchChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Point2_Pre_Ptch} placeholder="Enter text" onChange={this.point2PrePtchChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Point3_Pre_Ptch} placeholder="Enter text" onChange={this.point3PrePtchChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Point4_Pre_Ptch} placeholder="Enter text" onChange={this.point4PrePtchChange.bind(this)} />
                      </FormGroup>
                    </td>
                  </tr>
                  <tr>
                    <td>Roll</td>
                    <td>                  
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Point1_Pre_Roll} placeholder="Enter text" onChange={this.point1PreRollChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Point2_Pre_Roll} placeholder="Enter text" onChange={this.point2PreRollChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Point3_Pre_Roll} placeholder="Enter text" onChange={this.point3PreRollChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Point4_Pre_Roll} placeholder="Enter text" onChange={this.point4PreRollChange.bind(this)} />
                      </FormGroup>
                    </td>
                  </tr>
                </tbody>
              </Table>
            </Col>
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
