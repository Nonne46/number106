import React from "react";
import { createRoot } from "react-dom/client";
import { Toaster } from "react-hot-toast";

import "./style.css";
import App from "./App";

const container = document.getElementById("root");

const root = createRoot(container!);

root.render(
    <React.StrictMode>
        <App />
        <Toaster />
    </React.StrictMode>
);
