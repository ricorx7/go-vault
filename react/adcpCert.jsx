import React from 'react';
import ReactDOM from 'react-dom';
import {DataTable} from 'react-data-components';
import {blueGrey500} from 'material-ui/styles/colors';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import Toggle from 'material-ui/toggle';
import { Checkbox } from 'react-bootstrap';
import { Router, Route, Link, browserHistory } from 'react-router';
import { Button, Row, Col, Table, Glyphicon, FormControl } from 'react-bootstrap';

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


// List all the ADCP using the react-data-components.
export default class AdcpCert extends React.Component {
  
  constructor(props) {
    super(props);
    this.state = {
      data: {},
      tankTest: [],
      snrTest: [],
      waterTest: [],
    }
  }


  // At startup get all the ADCP data
  componentWillMount() {
    //this.serverRequest = this.loadAdcpFromServer();

    var urlSelected = "/vault/adcp/cert/"  + this.props.params.id;
    $.ajax({
      url: urlSelected,
      dataType: 'json',
      cache: false,
      success: function(data) {
        console.log("data length %i\n", data.length);
        this.setState({data: data});
      }.bind(this),
      error: function(xhr, status, err) {
        console.error("/vault/adcp", status, err.toString());
      }.bind(this)
    });

    console.log("data length: %i\n", this.state.data.length);
    console.log("adcp data: %i\n", this.state.data.length);
  }

  // Shutdown the display
  componentWillUnmount() {
      this.serverRequest.abort();
  }

  // Get the ADCP data from the database using AJAX
  loadAdcpFromServer() {
    var urlSelected = "/vault/adcp/cert/"  + this.props.params.id;
    $.ajax({
      url: urlSelected,
      dataType: 'json',
      cache: false,
      success: function(data) {
        console.log("data length %i\n", data.length);
        this.setState({data: data});
      }.bind(this),
      error: function(xhr, status, err) {
        console.error("/vault/adcp", status, err.toString());
      }.bind(this)
    });
  }

  // Render function
  render() {

      if(this.state.data.Adcp == null)
      {
          return(<div></div>);
      }

      var marginBottom = {
          bottomMargin: "20px"
      };

      var noBottom = {
          bottomMargin: "0px"
      };

      var marginLeft = {
          leftMargin: "5px"
      };

      var sigStyle = {
          border: "1px", 
          solid: "black", 
          marginTop: "5px"
      }

      var rowHeaderStyle = {
          bgcolor: "lightgray"
      }
    

    if(this.state.data.Adcp != null) 
    {
        var pressure;
        if(this.state.data.Adcp.PressureSensorPresent) {
            pressure = (<td><Glyphicon glyph="check" aria-hidden="true"></Glyphicon></td>)
        }
        var temperature;
        if (this.state.data.Adcp.TemperaturePresent) {
            temperature = (<td><Glyphicon glyph="check" aria-hidden="true"></Glyphicon></td>)
        }
        var recorderFormatted;
        if (this.state.data.Adcp.RecorderFormatted) {
            recorderFormatted = (<td><Glyphicon glyph="check" aria-hidden="true"></Glyphicon></td>)
        }
    }

    // Compass Cal Data
    if(this.state.data.CompassCal != null) 
    {    
        var compassCalData = [];
        this.state.data.CompassCal.map(function(cc, i) {
            var pt1 = Math.round(cc.Point1_Post_Hdg * 100) / 100;
            var pt1Diff = Math.round((cc.CompasscalBeam1Error) * 100) / 100;
            compassCalData.push(
                <tr>
                    <td><strong>0°</strong></td>
                    <td>{pt1}°</td>
                    <td>{pt1Diff}°</td>
                </tr>
            );
            var pt2 = Math.round(cc.Point2_Post_Hdg * 100) / 100;
            var pt2Diff = Math.round(cc.CompasscalBeam2Error * 100) / 100;
            compassCalData.push(
                <tr>
                    <td><strong>90°</strong></td>
                    <td>{pt2}°</td>
                    <td>{pt2Diff}°</td>
                </tr>
            );
            var pt3 = Math.round(cc.Point3_Post_Hdg * 100) / 100;
            var pt3Diff = Math.round(cc.CompasscalBeam3Error * 100) / 100;
            compassCalData.push(
                <tr>
                    <td><strong>180°</strong></td>
                    <td>{pt3}°</td>
                    <td>{pt3Diff}°</td>
                </tr>
            );
            var pt4 = Math.round(cc.Point4_Post_Hdg * 100) / 100;
            var pt4Diff = Math.round(cc.CompasscalBeam4Error * 100) / 100;
            compassCalData.push(
                <tr>
                    <td><strong>270°</strong></td>
                    <td>{pt4}°</td>
                    <td>{pt4Diff}°</td>
                </tr>
            );
        })
    }

    // Tank Test
    if(this.state.data.TankTest != null)
    {
        var tankTestData = [];
        this.state.data.TankTest.map(function(tt, i) {
        tankTestData.push(
        <tr>
            <td colSpan="5">{tt.SubsystemDescStr}</td>
        </tr>
        );
        var b0Noise = Math.round(tt.Beam0NoiseFloor * 100) / 100;
        var b0Sig = Math.round(tt.Beam0SignalTank * 100) / 100;
        tankTestData.push(
        <tr>
            <td><strong>Beam 0</strong></td>
            <td><Glyphicon glyph="check"></Glyphicon></td>
            <td>{b0Noise}</td>
            <td>{b0Sig}</td>
            <td><Glyphicon glyph="check"></Glyphicon></td>
        </tr>
        );
        var b1Noise = Math.round(tt.Beam1NoiseFloor * 100) / 100;
        var b1Sig = Math.round(tt.Beam1SignalTank * 100) / 100;
        tankTestData.push(
        <tr>
            <td><strong>Beam 1</strong></td>
            <td><Glyphicon glyph="check"></Glyphicon></td>
            <td>{b1Noise}</td>
            <td>{b1Sig}</td>
            <td><Glyphicon glyph="check"></Glyphicon></td>
        </tr>
        );
        var b2Noise = Math.round(tt.Beam2NoiseFloor * 100) / 100;
        var b2Sig = Math.round(tt.Beam2SignalTank * 100) / 100;
        tankTestData.push(
        <tr>
            <td><strong>Beam 2</strong></td>
            <td><Glyphicon glyph="check"></Glyphicon></td>
            <td>{b2Noise}</td>
            <td>{b2Sig}</td>
            <td><Glyphicon glyph="check"></Glyphicon></td>
        </tr>
        );
        var b3Noise = Math.round(tt.Beam3NoiseFloor * 100) / 100;
        var b3Sig = Math.round(tt.Beam3SignalTank * 100) / 100;
        tankTestData.push(
        <tr>
            <td><strong>Beam 3</strong></td>
            <td><Glyphicon glyph="check"></Glyphicon></td>
            <td>{b3Noise}</td>
            <td>{b3Sig}</td>
            <td><Glyphicon glyph="check"></Glyphicon></td>
        </tr>
        );
        })
    }

    // SNR test
    if(this.state.data.SnrTest != null)
    {
        var snrTestData = [];
        this.state.data.SnrTest.map(function(snr, i) {
        snrTestData.push(
            <tr style={rowHeaderStyle}>
                <td colSpan="5">{snr.SubsystemDescStr}</td>
            </tr>
        );
        var b0Sig = Math.round(snr.Beam0SignalLake * 100) / 100;
        var b1Sig = Math.round(snr.Beam1SignalLake * 100) / 100;
        var b2Sig = Math.round(snr.Beam2SignalLake * 100) / 100;
        var b3Sig = Math.round(snr.Beam3SignalLake * 100) / 100;
        snrTestData.push(
            <tr>
                <td><strong>Signal</strong></td>
                <td>{b0Sig}</td>
                <td>{b1Sig}</td>
                <td>{b2Sig}</td>
                <td>{b3Sig}</td>
            </tr>
        );
        var b0Noise = Math.round(snr.Beam0NoiseFloor * 100) / 100;
        var b1Noise = Math.round(snr.Beam1NoiseFloor * 100) / 100;
        var b2Noise = Math.round(snr.Beam2NoiseFloor * 100) / 100;
        var b3Noise = Math.round(snr.Beam3NoiseFloor * 100) / 100;
        snrTestData.push(
            <tr>
                <td><strong>Noise</strong></td>
                <td>{b0Noise}</td>
                <td>{b1Noise}</td>
                <td>{b2Noise}</td>
                <td>{b3Noise}</td>
            </tr>
        );
        var b0Snr = Math.round(snr.Beam0SnrLake * 100) / 100;
        var b1Snr = Math.round(snr.Beam1SnrLake * 100) / 100;
        var b2Snr = Math.round(snr.Beam2SnrLake * 100) / 100;
        var b3Snr = Math.round(snr.Beam3SnrLake * 100) / 100;
        snrTestData.push(
            <tr>
                <td><strong>SNR</strong></td>
                <td><b>{b0Snr}</b></td>
                <td><b>{b1Snr}</b></td>
                <td><b>{b2Snr}</b></td>
                <td><b>{b3Snr}</b></td>
            </tr>
        );
        })
    }

    // DMG Water test
    if(this.state.data.WaterTest != null)
    {
        var dmgTestData = [];
        this.state.data.WaterTest.map(function(wt, i) {
            dmgTestData.push(
                <tr>
                <td><strong>Beam Forward: {wt.TestOrientation}</strong></td>
                <td colSpan="4">{wt.SubsystemDescStr}</td>
                </tr>
            );
            dmgTestData.push(
                    <tr>
                    <td><strong>GPS</strong></td>
                    <td>{wt.GpsDistance}</td>
                    <td>{wt.GpsDirection}</td>
                    <td></td>
                    <td></td>
                    </tr>
            );
            dmgTestData.push(
                    <tr>
                    <td><strong>BT ENU</strong></td>
                    <td>{wt.BtDistance}</td>
                    <td>{wt.BtDirection}</td>
                    <td>{wt.DistanceError}</td>
                    <td>{wt.DirectionError}</td>
                    </tr>
            );
            })
        }

    return (
    <div visible-print-block>
		<Row>
			<Col xs={3}>
				<img alt="RTI" class="img-responsive text-center" src="/images/companylogo.png" />
			</Col>
			<Col xs={8}>
				<h2> ADCP/DVL Test Certificate </h2>
			</Col>
		</Row>
		<Row style={marginBottom}>
            <Col xs={5}>
                <dt><strong>Serial Number:</strong></dt><dd>{this.state.data.Adcp.SerialNumber}</dd>
            </Col>
            <Col xs={3}>
				<dt><strong>Order Number:</strong></dt><dd>{this.state.data.Adcp.OrderNumber}</dd>
			</Col>
            <Col xs={2}>
                <dt><strong>Customer:</strong></dt><dd>{this.state.data.Adcp.Customer}</dd>
            </Col>
		</Row>

		<Row style={marginBottom}>
			<Col xs={9}>
					<dt><strong>System:</strong></dt><dd>{this.state.data.Adcp.Frequency}</dd>
			</Col>
		</Row>

		<Row fluid>
			<Col xs={3}>
				<Table condensed style={noBottom}>
					<tbody>
						<tr>
							<td><strong>System Type:</strong></td>
							<td>{this.state.data.Adcp.SystemType}</td>
						</tr>
						<tr>
							<td><strong>Application:</strong></td>
							<td>{this.state.data.Adcp.Application}</td>
						</tr>
						<tr>
							<td><strong>Depth Rating:</strong></td>
							<td>{this.state.data.Adcp.DepthRating}</td>
						</tr>
                        <tr style={noBottom}>
                            <td><strong>Head Type:</strong></td>
                            <td>{this.state.data.Adcp.HeadType}</td>
                        </tr>
                        <tr style={noBottom}>
                            <td><strong>Hardware:</strong></td>
                            <td>{this.state.data.Adcp.Hardware}</td>
                        </tr>
                        <tr style={noBottom}>
                            <td><strong>Connector:</strong></td>
                            <td>{this.state.data.Adcp.ConnectorType}</td>
                        </tr>
                        <tr style={noBottom}>
                            <td><strong>Firmware:</strong></td>
                            <td>{this.state.data.Adcp.Firmware}</td>
                        </tr>
                        <tr style={noBottom}>
                            <td><strong>Software:</strong></td>
                            <td>{this.state.data.Adcp.Software}</td>
                        </tr>
                        <tr style={noBottom}>
                            <td><strong>Recorder Size:</strong></td>
                            <td>{this.state.data.Adcp.RecorderSize}</td>
                        </tr>
                        <tr style={noBottom}>
                            <td><strong>Recorder Formatted:</strong></td>
                            {recorderFormatted}
                        </tr>
                        <tr style={noBottom}>
                            <td><strong>Temperature:</strong></td>
                            {temperature}
                        </tr>
                        <tr style={noBottom}>
                            <td><strong>Pressure Sensor:</strong></td>
                            {pressure}
                        </tr>
                    </tbody>
                </Table>
                </Col>

                <Col xs={7} >
                    <dl>
                    <dt>
                        <strong>Compass Calibration - Heading Pitch Roll</strong>
                    </dt>
                    <dd>
                        <Table bordered condensed style={noBottom}>
                            <thead>
                                <tr>
                                    <th><h5><strong>Heading</strong></h5></th>
                                    <th><h5><strong>Reading (°)</strong></h5></th>
                                    <th><h5><strong>Error (°)</strong></h5></th>
                                </tr>
                            </thead>

                            <tbody>
                            {compassCalData}
						    </tbody>
					    </Table>
				    </dd>

                    <dt>
                        <strong>* ±2° error acceptance criteria</strong>
                    </dt>
                    <dt> <br /> </dt>

				<dt><strong>Beam Check - Tank Test</strong></dt>
				<dd>
					<Table bordered condensed>
						<thead>
							<tr>
								<th></th>
								<th><h5><strong>Correct Order</strong></h5></th>
								<th><h5><strong>Noise Floor</strong></h5></th>
								<th><h5><strong>Amplitude Tank</strong></h5></th>
								<th><h5><strong>Range OK</strong></h5></th>
							</tr>
						</thead>

						<tbody>
                        {tankTestData}
						</tbody>
					</Table>
				</dd>

				<dt><strong>Signal to Noise Ratio</strong></dt>
				<dd>
					<Table bordered condensed>
						<thead>
							<tr>
								<th></th>
								<th><h5><strong>Beam 0</strong></h5></th>
								<th><h5><strong>Beam 1</strong></h5></th>
								<th><h5><strong>Beam 2</strong></h5></th>
								<th><h5><strong>Beam 3</strong></h5></th>
							</tr>
						</thead>

						<tbody>
                        {snrTestData}
						</tbody>
					</Table>
				</dd>

			</dl>
		</Col>
	</Row>

    <Row fluid>
    <br />
    <br />
    <br />
    <br />
    <br />
    </Row>

    <Row fluid>
    <Col xs={10}>
        <dl>
        <dt>
            <strong>Distance Made Good</strong>
        </dt>
        <dd>
            <Table bordered condensed style={noBottom}>
            <thead>
                <tr>
                <th><h5><strong></strong></h5></th>
                <th><h5><strong>Distance (m)</strong></h5></th>
                <th><h5><strong>Direction (°)</strong></h5></th>
                <th><h5><strong>Distance Error (%)</strong></h5></th>
                <th><h5><strong>Direction Error (%)</strong></h5></th>
                </tr>
            </thead>

            <tbody>
            {dmgTestData}
            </tbody>
            </Table>
        </dd>
        </dl>
    </Col>
    </Row>

		<Row fluid style={marginLeft}>
			<Col xs={3} style={sigStyle}>
			    <h5>Tech Signature:</h5>
			</Col>
            <Col xs={5}>
            <FormControl type="text" />
            </Col>

			<Col xs={1} style={sigStyle}>
			    <h5>Date:</h5>
			</Col>
            <Col xs={3}>
                <FormControl type="text" />
            </Col>

		</Row>

		<Row fluid style={marginLeft}>
			<Col xs={3} style={sigStyle}>
			    <h5>QA/QC Signature:</h5>
			</Col>
            <Col xs={5}>
            <FormControl type="text" />
            </Col>
			<Col xs={1}>
			    <h5>Date:</h5>
			</Col>
            <Col xs={3}>
                <FormControl type="text" />
            </Col>

		</Row>

		<Row fluid style={{margin: "5px"}}>
			<Col xs={9}>
		    <small>
		    <b>Rowe Technology, Inc.</b> | 12655 Danielson Ct., Suite 306, Poway, California, USA<br />
		    http://www.rowetechinc.com | service@rowetechinc.com | +1 858 842 3020<br />
		    </small>
			</Col>
	    </Row>
	</div>

    );
  }
}

