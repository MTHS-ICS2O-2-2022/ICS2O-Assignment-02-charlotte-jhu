// Copyright (c) 2023 Charlote Jhu All rights reserved
//
// Created by: Charlotte Jhu
// Created on: March 2023
// This file contains the JS functions for index.html

'use strict'
/**
 * This function calculates the circumference of a circle
 */
function myButtonClicked() {
  // input
  const radius = parseInt(document.getElementById("radius").value)

  // process
  const π = 3.14
  const circumference = 2 * π * radius

  // output
  document.getElementById("circumference").innerHTML = "The circumference is " + circumference + "cm"
}
