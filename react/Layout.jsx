import React from "react";


import Navigation from "./header.jsx";

export default class Layout extends React.Component {
  render() {
    const { location } = this.props;
    const containerStyle = {
      marginTop: "10px"
    };
    console.log("layout");
    return (
      <div>
        <Navigation />

        <div class="container" style={containerStyle}>
          <div class="row">
            <div class="col-lg-12">

              {this.props.children}

            </div>
          </div>
          
        </div>
      </div>

    );
  }
}