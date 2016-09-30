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


// Edit the Tank test data.
export default class TankTestEdit extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
        data: {},
    }
  }

    // At startup get all the Tank Test data
  componentDidMount() {
    this.apiGetTtSelected();
    console.log("data length %i\n", this.state.data);
  }

    // Call API to set IsSelect selection
    apiGetTtSelected() {
      var urlSelected = "/vault/tt/edit/" + this.props.params.id;
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
    apiPostTt(wtData) {
      var urlSelected = "/vault/tt/edit/" + this.state.data.id;
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
        this.apiPostTt(JSON.stringify(this.state.data));
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

    // Update the distance
	  gpsDistanceChange(e) {
        this.state.data.GpsDistance = e.target.value;             // Update the object
        this.update();                                            // Update DB and display 
    }

    // Update Direction
    gpsDirectionChange(e) {
        this.state.data.GpsDirection = e.target.value;            // Update the object
        this.update();                                            // Update DB and display 
    }

    // ADCP Distance
    btDistanceChange(e) {
        this.state.data.BtDistance = e.target.value;              // Update the object
        this.update();                                            // Update DB and display 
    }

    // ADCP Distance Error.
    distanceErrorChange(e) {
        this.state.data.DistanceError = e.target.value;           // Update the object
        this.update();                                            // Update DB and display 
    }

    // ADCP Direction
    btDirectionChange(e) {
        this.state.data.BtDirection = e.target.value;             // Update the object
        this.update();                                            // Update DB and display 
    }

    // ADCP Direction Error.
    directionErrorChange(e) {
        this.state.data.DirectionError = e.target.value;          // Update the object
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

    beam0SignalLakeChange(e) {
        this.state.data.Beam0SignalLake = parseFloat(e.target.value);  // Update the object
        this.update();                                             // Update DB and display   
    }

    beam1SignalLakeChange(e) {
        this.state.data.Beam1SignalLake = parseFloat(e.target.value);  // Update the object
        this.update();                                             // Update DB and display   
    }

    beam2SignalLakeChange(e) {
        this.state.data.Beam2SignalLake = parseFloat(e.target.value);  // Update the object
        this.update();                                             // Update DB and display   
    }

    beam3SignalLakeChange(e) {
        this.state.data.Beam3SignalLake = parseFloat(e.target.value);  // Update the object
        this.update();                                             // Update DB and display   
    }

    beam0SignalOceanChange(e) {
        this.state.data.Beam0SignalOcean = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    beam1SignalOceanChange(e) {
        this.state.data.Beam1SignalOcean = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    beam2SignalOceanChange(e) {
        this.state.data.Beam2SignalOcean = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    beam3SignalOceanChange(e) {
        this.state.data.Beam3SignalOcean = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    beam0NoiseFloorChange(e) {
        this.state.data.Beam0NoiseFloor = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    beam1NoiseFloorChange(e) {
        this.state.data.Beam1NoiseFloor = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    beam2NoiseFloorChange(e) {
        this.state.data.Beam2NoiseFloor = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    beam3NoiseFloorChange(e) {
        this.state.data.Beam3NoiseFloor = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    beam0SnrLakeChange(e) {
        this.state.data.Beam0SnrLake = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    beam1SnrLakeChange(e) {
        this.state.data.Beam1SnrLake = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    beam2SnrLakeChange(e) {
        this.state.data.Beam2SnrLake = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    beam3SnrLakeChange(e) {
        this.state.data.Beam3SnrLake = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    beam0SnrOceanChange(e) {
        this.state.data.Beam0SnrOcean = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    beam1SnrOceanChange(e) {
        this.state.data.Beam1SnrOcean = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    beam2SnrOceanChange(e) {
        this.state.data.Beam2SnrOcean = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    beam3SnrOceanChange(e) {
        this.state.data.Beam3SnrOcean = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    profileRangeBeam0Change(e) {
        this.state.data.ProfileRangeBeam0 = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    profileRangeBeam1Change(e) {
        this.state.data.ProfileRangeBeam1 = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    profileRangeBeam2Change(e) {
        this.state.data.ProfileRangeBeam2 = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    profileRangeBeam3Change(e) {
        this.state.data.ProfileRangeBeam3 = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    glitchCountBeam0Change(e) {
        this.state.data.GlitchCountBeam0 = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    glitchCountBeam1Change(e) {
        this.state.data.GlitchCountBeam1 = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    glitchCountBeam2Change(e) {
        this.state.data.GlitchCountBeam2 = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    glitchCountBeam3Change(e) {
        this.state.data.GlitchCountBeam3 = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    bottomTrackAmplitudeBeam0Change(e) {
        this.state.data.BottomTrackAmplitudeBeam0 = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    bottomTrackAmplitudeBeam1Change(e) {
        this.state.data.BottomTrackAmplitudeBeam1 = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    bottomTrackAmplitudeBeam2Change(e) {
        this.state.data.BottomTrackAmplitudeBeam2 = parseFloat(e.target.value);  // Update the object
        this.update();                                              // Update DB and display   
    }

    bottomTrackAmplitudeBeam3Change(e) {
        this.state.data.BottomTrackAmplitudeBeam3 = parseFloat(e.target.value);  // Update the object
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
            <Col xs={2}>
              <MuiThemeProvider muiTheme={muiTheme}>
                <div style={styles.block}>
                    <Toggle label="Is Selected:" defaultToggled={this.convertToBool(this.state.data.IsSelected)} onToggle={this.isSelectedChange.bind(this, this.state.data.IsSelected)} style={styles.toggle} />
                </div>
              </MuiThemeProvider>
            </Col>
          </Row>

          <Row>
            <Col xs={2}>
              <FormGroup controlId="formControlsSelect">
                <ControlLabel>Orientation:</ControlLabel>
                <FormControl componentClass="select" placeholder="Orientation" value={this.state.data.TestOrientation} onChange={this.orientationChange.bind(this)}>
                  <option value="0">0</option>
                  <option value="3">3</option>
                </FormControl>
              </FormGroup>
            </Col>
          </Row>

          <Row>
            <Col md={5}>
              <Panel header="GPS">
                <Col xs={5}>
                  <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                    <ControlLabel>GPS Distance:</ControlLabel>
                    <FormControl type="text" value={this.state.data.GpsDistance} placeholder="Enter text" onChange={this.gpsDistanceChange.bind(this)} />
                    <FormControl.Feedback /> 
                  </FormGroup>
                </Col>
                <Col xs={5}>
                  <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                    <ControlLabel>GPS Direction:</ControlLabel>
                    <FormControl type="text" value={this.state.data.GpsDirection} placeholder="Enter text" onChange={this.gpsDirectionChange.bind(this)} />
                    <FormControl.Feedback /> 
                  </FormGroup>
                </Col>
              </Panel>
            </Col>

            <Col md={5}>
              <Panel header="ADCP">
                <Row>
                    <Col xs={5}>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <ControlLabel>Distance Traveled:</ControlLabel>
                        <FormControl type="text" value={this.state.data.BtDistance} placeholder="Enter text" onChange={this.btDistanceChange.bind(this)} />
                        <FormControl.Feedback /> 
                      </FormGroup>
                    </Col>
                    <Col xs={5}>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <ControlLabel>Direction:</ControlLabel>
                        <FormControl type="text" value={this.state.data.BtDirection} placeholder="Enter text" onChange={this.btDirectionChange.bind(this)} />
                        <FormControl.Feedback /> 
                      </FormGroup>
                    </Col>
                </Row>

                <Row>
                    <Col xs={5}>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <ControlLabel>Distance Error:</ControlLabel>
                        <FormControl type="text" value={this.state.data.DistanceError} placeholder="Enter text" onChange={this.distanceErrorChange.bind(this)} />
                        <FormControl.Feedback /> 
                      </FormGroup>
                    </Col>
                    <Col xs={5}>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <ControlLabel>Direction Error:</ControlLabel>
                        <FormControl type="text" value={this.state.data.DirectionError} placeholder="Enter text" onChange={this.directionErrorChange.bind(this)} />
                        <FormControl.Feedback /> 
                      </FormGroup>
                    </Col>
                </Row>
              </Panel>
            </Col>
          </Row>

          <Row>
            <Col xs={10}>
              <Table striped bordered condensed responsive hover>
                <thead>
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
                    <td>Signal Lake</td>
                    <td>                  
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Beam0SignalLake} placeholder="Enter text" onChange={this.beam0SignalLakeChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Beam1SignalLake} placeholder="Enter text" onChange={this.beam1SignalLakeChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Beam2SignalLake} placeholder="Enter text" onChange={this.beam2SignalLakeChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Beam3SignalLake} placeholder="Enter text" onChange={this.beam3SignalLakeChange.bind(this)} />
                      </FormGroup>
                    </td>
                  </tr>
                  <tr>
                    <td>Signal Ocean</td>
                    <td>                  
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Beam0SignalOcean} placeholder="Enter text" onChange={this.beam0SignalOceanChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Beam1SignalOcean} placeholder="Enter text" onChange={this.beam1SignalOceanChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Beam2SignalOcean} placeholder="Enter text" onChange={this.beam2SignalOceanChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Beam3SignalOcean} placeholder="Enter text" onChange={this.beam3SignalOceanChange.bind(this)} />
                      </FormGroup>
                    </td>
                  </tr>
                  <tr>
                    <td>Noise Floor</td>
                    <td>                  
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Beam0NoiseFloor} placeholder="Enter text" onChange={this.beam0NoiseFloorChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Beam1NoiseFloor} placeholder="Enter text" onChange={this.beam1NoiseFloorChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Beam2NoiseFloor} placeholder="Enter text" onChange={this.beam2NoiseFloorChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Beam3NoiseFloor} placeholder="Enter text" onChange={this.beam3NoiseFloorChange.bind(this)} />
                      </FormGroup>
                    </td>
                  </tr>
                  <tr>
                    <td>SNR Lake</td>
                    <td>                  
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Beam0SnrLake} placeholder="Enter text" onChange={this.beam0SnrLakeChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Beam1SnrLake} placeholder="Enter text" onChange={this.beam1SnrLakeChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Beam2SnrLake} placeholder="Enter text" onChange={this.beam2SnrLakeChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Beam3SnrLake} placeholder="Enter text" onChange={this.beam3SnrLakeChange.bind(this)} />
                      </FormGroup>
                    </td>
                  </tr>
                  <tr>
                    <td>SNR Ocean</td>
                    <td>                  
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Beam0SnrOcean} placeholder="Enter text" onChange={this.beam0SnrOceanChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Beam1SnrOcean} placeholder="Enter text" onChange={this.beam1SnrOceanChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Beam2SnrOcean} placeholder="Enter text" onChange={this.beam2SnrOceanChange.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.Beam3SnrOcean} placeholder="Enter text" onChange={this.beam3SnrOceanChange.bind(this)} />
                      </FormGroup>
                    </td>
                  </tr>
                  <tr>
                    <td>Profile Range</td>
                    <td>                  
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.ProfileRangeBeam0} placeholder="Enter text" onChange={this.profileRangeBeam0Change.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.ProfileRangeBeam1} placeholder="Enter text" onChange={this.profileRangeBeam1Change.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.ProfileRangeBeam2} placeholder="Enter text" onChange={this.profileRangeBeam2Change.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.ProfileRangeBeam3} placeholder="Enter text" onChange={this.profileRangeBeam3Change.bind(this)} />
                      </FormGroup>
                    </td>
                  </tr>
                  <tr>
                    <td>Glitch Count</td>
                    <td>                  
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.GlitchCountBeam0} placeholder="Enter text" onChange={this.glitchCountBeam0Change.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.GlitchCountBeam1} placeholder="Enter text" onChange={this.glitchCountBeam1Change.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.GlitchCountBeam2} placeholder="Enter text" onChange={this.glitchCountBeam2Change.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.GlitchCountBeam3} placeholder="Enter text" onChange={this.glitchCountBeam3Change.bind(this)} />
                      </FormGroup>
                    </td>
                  </tr>
                  <tr>
                    <td>BT Amp</td>
                    <td>                  
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.BottomTrackAmplitudeBeam0} placeholder="Enter text" onChange={this.bottomTrackAmplitudeBeam0Change.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.BottomTrackAmplitudeBeam1} placeholder="Enter text" onChange={this.bottomTrackAmplitudeBeam1Change.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.BottomTrackAmplitudeBeam2} placeholder="Enter text" onChange={this.bottomTrackAmplitudeBeam2Change.bind(this)} />
                      </FormGroup>
                    </td>
                    <td>
                      <FormGroup controlId="formBasicText" validationState={this.getValidationState()} >
                        <FormControl type="text" value={this.state.data.BottomTrackAmplitudeBeam3} placeholder="Enter text" onChange={this.bottomTrackAmplitudeBeam3Change.bind(this)} />
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
