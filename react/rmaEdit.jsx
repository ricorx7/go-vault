import React from 'react';
import ReactDOM from 'react-dom';
import {DataTable} from 'react-data-components';
import {blueGrey500} from 'material-ui/styles/colors';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import Toggle from 'material-ui/toggle';
import {Card, CardActions, CardHeader, CardMedia, CardTitle, CardText} from 'material-ui/Card';
import { Button, Row, Col, Glyphicon, Form, FormControl, FormGroup, ControlLabel, HelpBlock, Well, Checkbox, Panel  } from 'react-bootstrap';
import DatePicker from 'material-ui/DatePicker';
import TextField from 'material-ui/TextField';
import SelectField from 'material-ui/SelectField';
import MenuItem from 'material-ui/MenuItem';
import Divider from 'material-ui/Divider';
import RaisedButton from 'material-ui/RaisedButton';
import {Table, TableBody, TableHeader, TableHeaderColumn, TableRow, TableRowColumn} from 'material-ui/Table';
import Paper from 'material-ui/Paper';

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

const styles1 = {
  card: {
    //position: 'relative',
    //width: '350px',
    //color: 'red',
    //borderStyle: 'solid',
    //borderColor: 'yellowgreen'
    boxShadow: 'rgba(255, 0, 0, 0.117647) 0px 1px 16px, rgba(255, 0, 0, 0.117647) 0px 1px 14px',
    margin: '10px'
  },
  menu: {
    position: 'absolute',
    right: '10px',
    top: '15px'
  },
  cardHeader: {
    paddingBottom: '40px'
  }
}


  // Convert the date to "MM/DD/YYYY"
  //Use this method - it does handle double digits correctly
  Date.prototype.yyyymmdd = function() {
    var mm = (this.getMonth() + 1).toString(); // getMonth() is zero-based
    var dd = this.getDate().toString();

    return [ mm.length===2 ? '' : '0', mm, '/', dd.length===2 ? '' : '0', dd, '/', this.getFullYear(),].join(''); // padding
  };


// Edit the Tank test data.
export default class RmaEdit extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
        data: {},
        startDate: new Date(),
        receiveDate: new Date(),
        inspectionDate: new Date(),
        repairDate: new Date(),
        shipmentDate: new Date(),
        Products: [],
        ReceiveQty: "",
        ReceivePartNumber: "",
        ReceiveSerialNumber: "",
        RepairProducts: [],
        RepairQty: "",
        RepairPartNumber: "",
        RepairSerialNumber: "",
    }
  }

    // At startup get all the Tank Test data
  componentDidMount() {
    console.log(this.props.params.id);

    if(this.props.params.id == "add") {
      console.log("Add new RMA");
      
      // Find new RMA number
      this.apiAddRmaSelected();

    } else {
      console.log("Edit RMA");
      this.apiGetRmaSelected();
    }
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
          
          if(this.state.data.RmaDate) {
            this.setState({startDate: new Date(this.state.data.RmaDate)});
          } else {
            this.setState({startDate: new Date()});
          }

          if(this.state.data.ReceiveDate) {
            this.setState({receiveDate: new Date(this.state.data.ReceiveDate)});
          } else {
            this.setState({receiveDate: new Date()});
          }

          if(this.state.data.InspectionDate) {
            this.setState({inspectionDate: new Date(this.state.data.InspectionDate)});
          } else {
            this.setState({inspectionDate: new Date()});
          }

          if(this.state.data.RepairDate) {
            this.setState({repairDate: new Date(this.state.data.RepairDate)});
          } else {
            this.setState({repairDate: new Date()});
          }

          if(this.state.data.ShipDate) {
            this.setState({shipmentDate: new Date(this.state.data.ShipDate)});
          } else {
            this.setState({shipmentDate: new Date()});
          }
          
          this.setState({Products: this.state.data.Products});
          this.setState({RepairProducts: this.state.data.RepairProducts});
        }.bind(this),
        error: function(xhr, status, err) {
          console.error(urlSelected, status, err.toString());
        }.bind(this)
      });
    }

    // Call API to set IsSelect selection
    apiAddRmaSelected() {
      var urlSelected = "/vault/rma/add";
      $.ajax({
        url: urlSelected,
        dataType: 'json',
        cache: false,
        success: function(data) {
          console.log("Data gotten from %s\n", urlSelected);
          console.log(data);
          this.setState({data: data});
          this.setState({startDate: new Date(this.state.data.RmaDate)});
          this.setState({receiveDate: new Date(this.state.data.ReceiveDate)});
          this.setState({inspectionDate: new Date(this.state.data.InspectionDate)});
          this.setState({repairDate: new Date(this.state.data.RepairDate)});
          this.setState({shipmentDate: new Date(this.state.data.ShipDate)});
          this.setState({Products: this.state.data.Products});
          this.setState({RepairProducts: this.state.data.RepairProducts});
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
    rmaTypeChange(event, index, value) {
        this.state.data.RmaType = value.toString();
        this.update();                                              // Update DB and display   
    }

    // Set the Customer.
    companyChange(e) {
        this.state.data.Company = e.target.value;               // Update the object
        this.update();                                              // Update DB and display   
    }

    startDateChange(event, date) {
      this.setState({startDate: date});
      this.state.data.RmaDate = date.yyyymmdd();
      this.update();
    }

    contactNameChange(e) {
      this.state.data.ContactName = e.target.value;               // Update the object
      this.update();                                              // Update DB and display  
    }

    statusChange(event, index, value) {
      this.state.data.Status = value;
      this.update();
    }

    origSalesOrderChange(e) {
      this.state.data.OrigSalesOrder = e.target.value;
      this.update();
    }

    contactAddressChange(e) {
      this.state.data.ContactAddress = e.target.value;
      this.update();
    }

    contactAddress2Change(e) {
      this.state.data.ContactAddress2 = e.target.value;
      this.update();
    }

    contactCityStateZipChange(e) {
      this.state.data.ContactAddressCityStateZip = e.target.value;
      this.update();
    }

    contactCountryChange(e) {
      this.state.data.ContactAddressCountry = e.target.value;
      this.update();
    }

    contactEmailChange(e) {
      this.state.data.ContactEmail = e.target.value;
      this.update();
    }

    contactPhoneChange(e) {
      this.state.data.ContactPhone = e.target.value;
      this.update();
    }

    reasonReturnChange(e) {
      this.state.data.ReasonReturn = e.target.value;
      this.update();
    }

    returnCompanyChange(e) {
      this.state.data.ReturnCompany = e.target.value;
      this.update();
    }

    returnNameChange(e) {
      this.state.data.ReturnName = e.target.value;
      this.update();
    }

    returnAddressChange(e) {
      this.state.data.ReturnAddress = e.target.value;
      this.update();
    }

    returnAddress2Change(e) {
      this.state.data.ReturnAddressCont = e.target.value;
      this.update();
    }

    returnCityStateZipChange(e) {
      this.state.data.ReturnAddressCityStateZip = e.target.value;
      this.update();
    }

    returnCountryChange(e) {
      this.state.data.ReturnAddressCountry = e.target.value;
      this.update();
    }

    returnEmailChange(e) {
      this.state.data.ReturnEmail = e.target.value;
      this.update();
    }

    returnPhoneChange(e) {
      this.state.data.ReturnPhone = e.target.value;
      this.update();
    }

    sameAsContactChange(e) {
      this.state.data.ReturnCompany =  this.state.data.Company;
      this.state.data.ReturnContact = this.state.data.ContactName;
      this.state.data.ReturnAddress = this.state.data.ContactAddress;
      this.state.data.ReturnAddressCont = this.state.data.ContactAddress2;
      this.state.data.ReturnAddressCityStateZip = this.state.data.ContactAddressCityStateZip;
      this.state.data.ReturnAddressCountry = this.state.data.ContactAddressCountry;
      this.state.data.ReturnEmail = this.state.data.ContactEmail;
      this.state.data.ReturnPhone = this.state.data.ContactPhone;
      this.forceUpdate(); // Render display with updates
      this.update();
    }

    receiveDateChange(event, date) {
      this.setState({receiveDate: date});
      this.state.data.ReceiveDate = date.yyyymmdd();
      this.state.data.Status = "Received";
      this.update();
    }

    receiveUserChange(e) {
      this.state.data.ReceiveUser = e.target.value;
      this.update();
    }

    receiveInfoChange(e) {
      this.state.data.ReceiveInfo = e.target.value;
      this.update();
    }

    receiveQtyChange(e) {
      this.setState({ReceiveQty: e.target.value});
    }

    receivePartNumberChange(e) {
      this.setState({ReceivePartNumber: e.target.value});
    }

    receiveSerialNumberChange(e) {
      this.setState({ReceiveSerialNumber: e.target.value});
    }

    addProductChange(event){

        var product = {"PartNumber": this.state.ReceivePartNumber,
                      "SerialNumber": this.state.ReceiveSerialNumber,
                      "Qty": parseInt(this.state.ReceiveQty)
                      };

        var newArray = this.state.Products.slice();    
        newArray.push(product);   
        this.setState({Products:newArray})

        this.state.data.Products.push(product);
        this.update();

        // Clear the entries
        this.setState({ReceivePartNumber: ""});
        this.setState({ReceiveSerialNumber: ""});
        this.setState({ReceiveQty: ""});
    }

    inspectionDateChange(event, date) {
      this.setState({inspectionDate: date});
      this.state.data.InspectionDate = date.yyyymmdd();
      this.state.data.Status = "Inspected";
      this.update();
    }

    inspectionUserChange(e) {
      this.state.data.InspectionUser = e.target.value;
      this.update();
    }

    inspectionInfoChange(e) {
      this.state.data.InspectionInfo = e.target.value;
      this.update();
    }

    repairDateChange(event, date) {
      this.setState({repairDate: date});
      this.state.data.RepairDate = date.yyyymmdd();
      this.state.data.Status = "Repaired";
      this.update();
    }

    repairUserChange(e) {
      this.state.data.RepairUser = e.target.value;
      this.update();
    }

    repairInfoChange(e) {
      this.state.data.RepairInfo = e.target.value;
      this.update();
    }

    repairQtyChange(e) {
      this.setState({RepairQty: e.target.value});
    }

    repairPartNumberChange(e) {
      this.setState({RepairPartNumber: e.target.value});
    }

    repairSerialNumberChange(e) {
      this.setState({RepairSerialNumber: e.target.value});
    }

    addRepairProductChange(event){

        var product = {"PartNumber": this.state.RepairPartNumber,
                      "SerialNumber": this.state.RepairSerialNumber,
                      "Qty": parseInt(this.state.RepairQty)
                      };

        var newArray = this.state.RepairProducts.slice();    
        newArray.push(product);   
        this.setState({RepairProducts:newArray})

        this.state.data.RepairProducts.push(product);
        this.update();

        // Clear the entries
        this.setState({RepairPartNumber: ""});
        this.setState({RepairSerialNumber: ""});
        this.setState({RepairQty: ""});
    }

    estRepairHoursChange(e) {
        var hr = parseInt(e.target.value)
        if(isNaN(hr)) {
          this.state.data.RepairEstHours = 0;
          console.log("NaN");
        } else {
          this.state.data.RepairEstHours = hr;                    // Update the object
        }
        this.update();
    }

    billableChange(event, index, value) {
        this.state.data.Billable = value.toString();                    // Update the object
        this.update();
    }

    quoteNumChange(e) {
      this.state.data.QuoteNum = e.target.value;                    // Update the object
      this.update();
    }

    shipmentDateChange(event, date) {
      this.setState({shipmentDate: date});
      this.state.data.ShipDate = date.yyyymmdd();
      this.state.data.Status = "Returned";
      this.update();
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

    const paperStyle = {
      height: 80,
      width: 1000,
      margin: 10,
      textIndent: 20,
      textAlign: 'left',
      display: 'inline-block',
    };


    return (
        <div style={containerStyle}>
          <MuiThemeProvider muiTheme={muiTheme}>
            <Paper style={paperStyle} zDepth={5}>
              <h1>{this.state.data.RmaType}{this.state.data.RmaNumber} - {this.state.data.Company}</h1>
            </Paper>
          </MuiThemeProvider>
          
          <Row>
            <Col sm={12}>
            <MuiThemeProvider muiTheme={muiTheme}>
              <Card initiallyExpanded={true} style={styles1.card}>
                  <CardHeader title="RMA Info" subtitle="" actAsExpander={true} showExpandableButton={true} />
                  <CardText expandable={true}>


                <SelectField floatingLabelText="RMA Type" value={this.state.data.RmaType} onChange={this.rmaTypeChange.bind(this)}>
                  <MenuItem value={"290"} primaryText="290 - Warranty" />
                  <MenuItem value={"280"} primaryText="280 - Billable" />
                  <MenuItem value={"259"} primaryText="259 - Demo Repair" />
                </SelectField>
                <br />
                <TextField hintText="RMA Number" floatingLabelText="RMA Number" value={this.state.data.RmaNumber} onChange={this.rmaNumberChange.bind(this)} />
                <br />
                <DatePicker hintText="RMA date created" floatingLabelText="RMA date created" value={this.state.startDate} autoOk={true} locale="en-US" onChange={this.startDateChange.bind(this)} />
                <br />
                <SelectField floatingLabelText="Status" value={this.state.data.Status} onChange={this.statusChange.bind(this)}>
                  <MenuItem value={"Reported"} primaryText="Reported" />
                  <MenuItem value={"Received"} primaryText="Received" />
                  <MenuItem value={"Inspected"} primaryText="Inspected" />
                  <MenuItem value={"Repaired"} primaryText="Repaired" />
                  <MenuItem value={"Returned"} primaryText="Returned" />
                  <MenuItem value={"Completed"} primaryText="Completed" />
                </SelectField>
                <br />
                <TextField hintText="Original Salesorder" floatingLabelText="Original Salesorder" value={this.state.data.OrigSalesOrder} onChange={this.origSalesOrderChange.bind(this)} />
              </CardText>
            </Card>
            </MuiThemeProvider>
            </Col>
        </Row>

        <Row>
          <Col sm={6}>
          <MuiThemeProvider muiTheme={muiTheme}>
            <Card initiallyExpanded={true} style={styles1.card}>
              <CardHeader
                    title="Customer Contact Info"
                    subtitle=""
                    actAsExpander={true}
                    showExpandableButton={true}
                  />
                  <CardActions>
                  </CardActions>
                  <CardText expandable={true}>        
                    <TextField hintText="Company" floatingLabelText="Company" value={this.state.data.Company} onChange={this.companyChange.bind(this)} /> 
                    <br />
                    <TextField hintText="Contact Name" floatingLabelText="Contact Name" value={this.state.data.ContactName} onChange={this.contactNameChange.bind(this)} />
                    <br />
                    <TextField hintText="Address" floatingLabelText="Address" value={this.state.data.ContactAddress} onChange={this.contactAddressChange.bind(this)} />
                    <br />
                    <TextField hintText="Address Line 2" floatingLabelText="Address Line 2" value={this.state.data.ContactAddress2} onChange={this.contactAddress2Change.bind(this)} />
                    <br />
                    <TextField hintText="City, State, Zip" floatingLabelText="City, State, Zip" value={this.state.data.ContactAddressCityStateZip} onChange={this.contactCityStateZipChange.bind(this)} />
                    <br />
                    <TextField hintText="Country" floatingLabelText="Country" value={this.state.data.ContactAddressCountry} onChange={this.contactCountryChange.bind(this)} />
                    <br />
                    <TextField hintText="Contact Email" floatingLabelText="Contact Email" value={this.state.data.ContactEmail} onChange={this.contactEmailChange.bind(this)} />
                    <br />
                    <TextField hintText="Contact Phone" floatingLabelText="Contact Phone" value={this.state.data.ContactPhone} onChange={this.contactPhoneChange.bind(this)} />
                    <br />
                    <TextField hintText="Reason for Return" floatingLabelText="Reason for Return" fullWidth={true} multiLine={true} rows={5} rowsMax={10} value={this.state.data.ReasonReturn} onChange={this.reasonReturnChange.bind(this)} />
                  </CardText>
              </Card>
            </MuiThemeProvider>
          </Col>

          <Col sm={6}>
          <MuiThemeProvider muiTheme={muiTheme}>
            <Card initiallyExpanded={true} style={styles1.card}>
              <CardHeader
                    title="Return Shipping Information"
                    subtitle=""
                    actAsExpander={true}
                    showExpandableButton={true}
                  />
                  <CardActions>
                  </CardActions>
                  <CardText expandable={true}> 
                    <form onSubmit={this.sameAsContactChange.bind(this)}>     
                      <RaisedButton label="Same as Contact" secondary={true} type="submit" />
                    </form>
                    <br />
                    <TextField hintText="Company" floatingLabelText="Company" value={this.state.data.ReturnCompany} onChange={this.returnCompanyChange.bind(this)} /> 
                    <br />
                    <TextField hintText="Contact Name" floatingLabelText="Contact Name" value={this.state.data.ReturnContact} onChange={this.returnNameChange.bind(this)} />
                    <br />
                    <TextField hintText="Address" floatingLabelText="Address" value={this.state.data.ReturnAddress} onChange={this.returnAddressChange.bind(this)} />
                    <br />
                    <TextField hintText="Address Line 2" floatingLabelText="Address Line 2" value={this.state.data.ReturnAddressCont} onChange={this.returnAddress2Change.bind(this)} />
                    <br />
                    <TextField hintText="City, State, Zip" floatingLabelText="City, State, Zip" value={this.state.data.ReturnAddressCityStateZip} onChange={this.returnCityStateZipChange.bind(this)} />
                    <br />
                    <TextField hintText="Country" floatingLabelText="Country" value={this.state.data.ReturnAddressCountry} onChange={this.returnCountryChange.bind(this)} />
                    <br />
                    <TextField hintText="Contact Email" floatingLabelText="Contact Email" value={this.state.data.ReturnEmail} onChange={this.returnEmailChange.bind(this)} />
                    <br />
                    <TextField hintText="Contact Phone" floatingLabelText="Contact Phone" value={this.state.data.ReturnPhone} onChange={this.returnPhoneChange.bind(this)} />
                  </CardText>
              </Card>
            </MuiThemeProvider>
          </Col>

        </Row>

          <Row>
            <Col sm={12}>
              <MuiThemeProvider muiTheme={muiTheme}>
                <Card initiallyExpanded={true} style={styles1.card}>
                  <CardHeader title="Receive Information" subtitle="" actAsExpander={true} showExpandableButton={true} />
                  <CardText expandable={true}>
                    <Row>
                    <Col sm={5}>
                      <DatePicker hintText="Date Received" floatingLabelText="Date Received" value={this.state.receiveDate} autoOk={true} locale="en-US" onChange={this.receiveDateChange.bind(this)} />
                      <br />
                      <TextField hintText="Received By" floatingLabelText="Received By" value={this.state.data.ReceiveUser} onChange={this.receiveUserChange.bind(this)} />
                    </Col>
                    <Col sm={5}>
                      <TextField hintText="Receive Information" floatingLabelText="Receive Information" fullWidth={true} multiLine={true} rows={5} rowsMax={10} value={this.state.data.ReceiveInfo} onChange={this.receiveInfoChange.bind(this)} />
                    </Col>
                  </Row>
                  <Row>
                  <Table height='200px' fixedHeader={true} fixedFooter={false} selectable={false} multiSelectable={true} >
                        <TableHeader displaySelectAll={false} adjustForCheckbox={false} enableSelectAll={true}>
                          <TableRow>
                            <TableHeaderColumn colSpan="4" tooltip="Received Products" style={{textAlign: 'center'}}>
                              Received Products
                            </TableHeaderColumn>
                          </TableRow>
                          <TableRow>
                            <TableHeaderColumn tooltip="ID">ID</TableHeaderColumn>
                            <TableHeaderColumn tooltip="Qty">Qty</TableHeaderColumn>
                            <TableHeaderColumn tooltip="Serial Number">Serial Number</TableHeaderColumn>
                            <TableHeaderColumn tooltip="Part Number">Part Number</TableHeaderColumn>
                          </TableRow>
                        </TableHeader>
                        <TableBody displayRowCheckbox={false} deselectOnClickaway={true} showRowHover={true} stripedRows={false}>
                          {this.state.Products.map( (row, index) => (
                            <TableRow key={index} selected={row.selected}>
                              <TableRowColumn>{index}</TableRowColumn>
                              <TableRowColumn>{row.Qty}</TableRowColumn>
                              <TableRowColumn>{row.SerialNumber}</TableRowColumn>
                              <TableRowColumn>{row.PartNumber}</TableRowColumn>
                            </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                    </Row>

                    <form onSubmit={this.addProductChange.bind(this)}> 
                      <TextField hintText="Qty" floatingLabelText="Qty" value={this.state.ReceiveQty} onChange={this.receiveQtyChange.bind(this)} style={containerStyle} />
                      <TextField hintText="Serial Number" floatingLabelText="Serial Number" value={this.state.ReceiveSerialNumber} onChange={this.receiveSerialNumberChange.bind(this)} style={containerStyle} />
                      <TextField hintText="Part Number" floatingLabelText="Part Number" value={this.state.ReceivePartNumber} onChange={this.receivePartNumberChange.bind(this)} style={containerStyle} />
                      <RaisedButton label="ADD" primary={true} type="submit" style={containerStyle} />
                    </form>
                  </CardText>
                </Card>
              </MuiThemeProvider>
            </Col>
          </Row>

          <Row>
            <Col sm={12}>
              <MuiThemeProvider muiTheme={muiTheme}>
                <Card initiallyExpanded={true} style={styles1.card}>
                  <CardHeader title="Inspection Information" subtitle="" actAsExpander={true} showExpandableButton={true} />
                  <CardText expandable={true}>
                    <Row>
                    <Col sm={5}>
                      <DatePicker hintText="Date Inspected" floatingLabelText="Date Inspected" value={this.state.inspectionDate} autoOk={true} locale="en-US" onChange={this.inspectionDateChange.bind(this)} />
                      <br />
                      <TextField hintText="Inspected By" floatingLabelText="Inspected By" value={this.state.data.InspectionUser} onChange={this.inspectionUserChange.bind(this)} />
                    </Col>
                    <Col sm={5}>
                      <TextField hintText="Inspection Information" floatingLabelText="Inspection Information" fullWidth={true} multiLine={true} rows={5} rowsMax={10} value={this.state.data.InspectionInfo} onChange={this.inspectionInfoChange.bind(this)} />
                    </Col>
                  </Row>
                  </CardText>
                </Card>
              </MuiThemeProvider>
            </Col>
          </Row>

          <Row>
            <Col sm={12}>
              <MuiThemeProvider muiTheme={muiTheme}>
                <Card initiallyExpanded={true} style={styles1.card}>
                  <CardHeader title="Repair Information" subtitle="" actAsExpander={true} showExpandableButton={true} />
                  <CardText expandable={true}>
                    <Row>
                    <Col sm={5}>
                      <DatePicker hintText="Date Repaired" floatingLabelText="Date Repaired" value={this.state.repairDate} autoOk={true} locale="en-US" onChange={this.repairDateChange.bind(this)} />
                      <br />
                      <TextField hintText="Repaired By" floatingLabelText="Repaired By" value={this.state.data.RepairUser} onChange={this.repairUserChange.bind(this)} />
                      <br />
                      <TextField hintText="Ext. Repair Hours" floatingLabelText="Est. Repair Hours" value={this.state.data.RepairEstHours} onChange={this.estRepairHoursChange.bind(this)} />
                      <br />
                      <SelectField floatingLabelText="Billable" value={this.state.data.Billable} onChange={this.billableChange.bind(this)}>
                        <MenuItem value={"Warranty"} primaryText="Warranty" />
                        <MenuItem value={"Billable"} primaryText="Billable" />
                        <MenuItem value={"N/A"} primaryText="N/A" />
                      </SelectField>
                      <br />
                      <TextField hintText="Quote Number" floatingLabelText="Quote Number" value={this.state.data.QuoteNum} onChange={this.quoteNumChange.bind(this)} />
                    </Col>
                    <Col sm={5}>
                      <TextField hintText="Repair Information" floatingLabelText="Repair Information" fullWidth={true} multiLine={true} rows={10} rowsMax={10} value={this.state.data.RepairInfo} onChange={this.repairInfoChange.bind(this)} />
                    </Col>
                  </Row>
                  <Row>
                  <Table height='200px' fixedHeader={true} fixedFooter={false} selectable={false} multiSelectable={true} >
                        <TableHeader displaySelectAll={false} adjustForCheckbox={false} enableSelectAll={true}>
                          <TableRow>
                            <TableHeaderColumn colSpan="4" tooltip="Repaired Products" style={{textAlign: 'center'}}>
                              Repaired Products
                            </TableHeaderColumn>
                          </TableRow>
                          <TableRow>
                            <TableHeaderColumn tooltip="ID">ID</TableHeaderColumn>
                            <TableHeaderColumn tooltip="Qty">Qty</TableHeaderColumn>
                            <TableHeaderColumn tooltip="Serial Number">Serial Number</TableHeaderColumn>
                            <TableHeaderColumn tooltip="Part Number">Part Number</TableHeaderColumn>
                          </TableRow>
                        </TableHeader>
                        <TableBody displayRowCheckbox={false} deselectOnClickaway={true} showRowHover={true} stripedRows={false}>
                          {this.state.RepairProducts.map( (row, index) => (
                            <TableRow key={index} selected={row.selected}>
                              <TableRowColumn>{index}</TableRowColumn>
                              <TableRowColumn>{row.Qty}</TableRowColumn>
                              <TableRowColumn>{row.SerialNumber}</TableRowColumn>
                              <TableRowColumn>{row.PartNumber}</TableRowColumn>
                            </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                    </Row>

                    <form onSubmit={this.addRepairProductChange.bind(this)}> 
                      <TextField hintText="Qty" floatingLabelText="Qty" value={this.state.RepairQty} onChange={this.repairQtyChange.bind(this)} style={containerStyle} />
                      <TextField hintText="Serial Number" floatingLabelText="Serial Number" value={this.state.RepairSerialNumber} onChange={this.repairSerialNumberChange.bind(this)} style={containerStyle} />
                      <TextField hintText="Part Number" floatingLabelText="Part Number" value={this.state.RepairPartNumber} onChange={this.repairPartNumberChange.bind(this)} style={containerStyle} />
                      <RaisedButton label="ADD" primary={true} type="submit" style={containerStyle} />
                    </form>
                  </CardText>
                </Card>
              </MuiThemeProvider>
            </Col>
          </Row>

          <Row>
            <Col sm={12}>
              <MuiThemeProvider muiTheme={muiTheme}>
                <Card initiallyExpanded={true} style={styles1.card}>
                  <CardHeader title="Shipment Information" subtitle="" actAsExpander={true} showExpandableButton={true} />
                  <CardText expandable={true}>
                    <Row>
                    <Col sm={5}>
                      <DatePicker hintText="Date Shipped Back" floatingLabelText="Date Shipped Back" value={this.state.shipmentDate} autoOk={true} locale="en-US" onChange={this.shipmentDateChange.bind(this)} />
                    </Col>
                  </Row>
                  </CardText>
                </Card>
              </MuiThemeProvider>
            </Col>
          </Row>

          <Row>
            <Col sm={12}>
              <MuiThemeProvider muiTheme={muiTheme}>
                <Card initiallyExpanded={true} style={styles1.card}>
                  <CardHeader title="Notes" subtitle="" actAsExpander={true} showExpandableButton={true} />
                  <CardText expandable={true}>
                    <TextField hintText="Notes" floatingLabelText="Notes" fullWidth={true} multiLine={true} rows={5} rowsMax={10} value={this.state.data.Notes} onChange={this.notesChange.bind(this)} />
                  </CardText>
                </Card>
              </MuiThemeProvider>
            </Col>
          </Row>
        
        
        </div>
    );
  }
}

// Use the url PROP to get the Water Test data
//ReactDOM.render((<WaterTestEdit url="/vault/wt/edit/" />), document.getElementById('watertestEdit'));
