import AppBar from "@material-ui/core/AppBar";
import Box from "@material-ui/core/Box";
import { makeStyles, useTheme } from "@material-ui/core/styles";
import Tab from "@material-ui/core/Tab";
import Tabs from "@material-ui/core/Tabs";
import Typography from "@material-ui/core/Typography";
import PropTypes from "prop-types";
import React, { useState } from "react";
import SwipeableViews from "react-swipeable-views";
import ProdStages from "./ProdStages";
import ViewAllProdStages from "./ViewAllProdStages";

function TabPanel(props) {
  const { children, value, index, ...other } = props;
  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`full-width-tabpanel-${index}`}
      aria-labelledby={`full-width-tab-${index}`}
      {...other}
    >
      {value === index && (
        <Box p={3}>
          <Typography>{children}</Typography>
        </Box>
      )}
    </div>
  );
}
TabPanel.propTypes = {
  children: PropTypes.node,
  index: PropTypes.any.isRequired,
  value: PropTypes.any.isRequired,
};
function a11yProps(index) {
  return {
    id: `full-width-tab-${index}`,
    "aria-controls": `full-width-tabpanel-${index}`,
  };
}
const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    backgroundColor: theme.palette.background.paper,
  },
}));

//my code here

const initialState = {
  name: "",
  status: true,
  machine: true,
  machinename: "",
  human: true,
  humancount: 0,
  costhours: 0,
};

const viewModesVIEW = "VIEW";
const viewModesNEW = "NEW";
const viewModesEDIT = "EDIT";
const viewModesTOBEEDIT = "ToBeEdited";

const ProdStagesContainer = () => {
  const classes = useStyles();
  const theme = useTheme();

  //my states
  const [value, setValue] = React.useState(0);
  const [viewMode, setViewMode] = useState(viewModesVIEW);
  const [currentProd, setCurrentProd] = useState(initialState);

  const handleChange = (event, newValue) => {
    setValue(newValue);
  };

  const handleChangeIndex = (index) => {
    setValue(index);
  };

  return (
    <div className={classes.root}>
      <AppBar position="static" color="default">
        <Tabs
          value={value}
          onChange={handleChange}
          indicatorColor="primary"
          textColor="primary"
          aria-label="full width tabs example"
        >
          <Tab label="VIEW MATERIALS" {...a11yProps(0)} />
          <Tab label="VIEW ALL RAW MATERIALS" {...a11yProps(1)} />
        </Tabs>
      </AppBar>
      <SwipeableViews
        axis={theme.direction === "rtl" ? "x-reverse" : "x"}
        index={value}
        onChangeIndex={handleChangeIndex}
      >
        <TabPanel value={value} index={0} dir={theme.direction}>
          <ProdStages
            currentProd={currentProd}
            setCurrentProd={setCurrentProd}
            initialState={initialState}
            viewMode={viewMode}
            setViewMode={setViewMode}
            viewModesVIEW={viewModesVIEW}
            viewModesEDIT={viewModesEDIT}
            viewModesNEW={viewModesNEW}
            viewModesTOBEEDIT={viewModesTOBEEDIT}
          />
        </TabPanel>
        <TabPanel value={value} index={1} dir={theme.direction}>
          <ViewAllProdStages
            changeView={setValue}
            viewModesTOBEEDIT={viewModesTOBEEDIT}
            currentProd={currentProd}
            setCurrentProd={setCurrentProd}
            setViewMode={setViewMode}
          />
        </TabPanel>
      </SwipeableViews>
    </div>
  );
};

export default ProdStagesContainer;
