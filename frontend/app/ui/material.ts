"use client";

import "@lit-labs/ssr-react/enable-lit-ssr.js";
import { MdOutlinedButton } from "@material/web/button/outlined-button.js";
import { MdOutlinedTextField } from "@material/web/textfield/outlined-text-field.js";
import { createComponent } from "@lit/react";
import React from "react";

export const OutlinedTextField = createComponent({
  tagName: "md-outlined-text-field",
  elementClass: MdOutlinedTextField,
  react: React,
  events: {
    onactivate: "activate",
    onchange: "change",
  },
});

export const OutlinedButton = createComponent({
  tagName: "md-outlined-button",
  elementClass: MdOutlinedButton,
  react: React,
  events: {
    onclick: "click",
  },
});
