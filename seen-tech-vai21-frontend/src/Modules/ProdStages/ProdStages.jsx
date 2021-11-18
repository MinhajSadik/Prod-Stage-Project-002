import {
  Button,
  FormControlLabel,
  Grid,
  Switch,
  TextField,
} from "@material-ui/core";
import Checkbox from "@material-ui/core/Checkbox";
import Collapse from "@material-ui/core/Collapse";
import IconButton from "@material-ui/core/IconButton";
import Paper from "@material-ui/core/Paper";
import { makeStyles } from "@material-ui/core/styles";
import CloseIcon from "@material-ui/icons/Close";
import Alert from "@material-ui/lab/Alert";
import Axios from "axios";
import React, { useState } from "react";
import { apiBaseUrl } from "../../config.json";

const useStyles = makeStyles((theme) => ({
  Button: {
    marginRight: theme.spacing(1),
  },
  Switch: {
    textAlign: "center",
  },
  Paper: {
    width: "100%",
    height: "100%",
    padding: theme.spacing(2),
    marginTop: theme.spacing(2),
  },
  Table: {
    marginTop: "40px",
  },
  TextField: {
    marginTop: theme.spacing(1),
  },
  MarginLeft: {
    marginLeft: theme.spacing(4),
  },
  MarginBottom: {
    marginBottom: theme.spacing(1),
  },
}));

const ProdStages = ({
  //recive from prodStageContainer
  currentProd,
  setCurrentProd,
  initialState,
  viewMode,
  setViewMode,
  viewModesVIEW,
  viewModesNEW,
  viewModesEDIT,
  viewModesTOBEEDIT,
}) => {
  const classes = useStyles();

  const [openSuccess, setOpenSuccess] = useState(false);
  const [openFailure, setOpenFailure] = useState(false);

  const handleCreateNewProdStage = async (data) => {
    await Axios({
      method: "POST",
      url: `${apiBaseUrl}/prodstages/new`,
      data: JSON.stringify(currentProd),
      headers: {
        "content-type": "application/json",
      },
    })
      .then((res) => {
        console.log(res.data);
        setOpenSuccess(true);
      })
      .catch((err) => {
        setOpenFailure(err.response.data);
      });
    clear();
  };

  const handleEditProdStage = async () => {
    await Axios({
      method: "PUT",
      url: `${apiBaseUrl}/prodstages/modify`,
      data: JSON.stringify(currentProd),
      headers: {
        "content-type": "application/json",
      },
    })
      .then((res) => {
        console.log(res.data.results);
      })
      .catch((err) => {
        console.log(err);
        setOpenFailure(err.response.data);
      });
    clear();
  };

  //State Clear Materials//
  const clear = () => {
    setCurrentProd(initialState);
    setViewMode(viewModesVIEW);
  };

  return (
    <>
      <Grid container>
        <Grid item xs={12} className={classes.root}>
          <Button
            style={{ marginRight: "10px" }}
            variant="contained"
            color="secondary"
            size="small"
            disabled={viewMode === viewModesEDIT}
            onClick={(e) => {
              setCurrentProd(initialState);
              setViewMode(viewModesNEW);
            }}
          >
            CREATE NEW
          </Button>

          <Button
            className={classes.Button}
            variant="contained"
            color="secondary"
            size="small"
            disabled={
              viewMode === viewModesNEW ||
              viewMode === viewModesVIEW ||
              viewMode === viewModesEDIT
            }
            onClick={(e) => {
              setViewMode(viewModesEDIT);
            }}
          >
            EDIT
          </Button>

          <Button
            variant="contained"
            color="primary"
            className={classes.Button}
            size="small"
            disabled={viewMode !== viewModesNEW && viewMode !== viewModesEDIT}
            onClick={(e) => {
              if (viewMode === viewModesNEW) {
                handleCreateNewProdStage();
              } else {
                handleEditProdStage();
              }
            }}
          >
            SAVE
          </Button>

          <Button
            variant="contained"
            className={classes.Button}
            size="small"
            onClick={(e) => {
              clear();
            }}
          >
            CANCEL
          </Button>
        </Grid>
        <Paper className={classes.Paper}>
          <Collapse in={openSuccess} timeout="auto">
            <Alert
              action={
                <IconButton
                  aria-label="close"
                  color="inherit"
                  size="small"
                  onClick={() => {
                    setOpenSuccess(false);
                  }}
                >
                  <CloseIcon fontSize="inherit" />
                </IconButton>
              }
            >
              Saved Successfully
            </Alert>
          </Collapse>

          <Collapse in={openFailure} timeout="auto">
            <Alert
              severity="error"
              action={
                <IconButton
                  aria-label="close"
                  color="inherit"
                  size="small"
                  onClick={() => {
                    setOpenFailure(false);
                  }}
                >
                  <CloseIcon fontSize="inherit" />
                </IconButton>
              }
            >
              {openFailure}
            </Alert>
          </Collapse>
          <Grid item xs={12} className={classes.Switch}>
            <FormControlLabel
              control={
                <Switch
                  checked={currentProd.status}
                  disabled={
                    viewMode !== viewModesNEW && viewMode !== viewModesEDIT
                  }
                  onChange={(e) =>
                    setCurrentProd({
                      ...currentProd,
                      status: e.target.checked,
                    })
                  }
                  name="checked"
                  color="primary"
                />
              }
              label="ACTIVE"
            />
          </Grid>
          <Grid container spacing={1} className={classes.TextField}>
            <Grid item xs={12}>
              <TextField
                size="small"
                value={currentProd.name}
                disabled={
                  viewMode !== viewModesNEW && viewMode !== viewModesEDIT
                }
                onChange={(e) => {
                  setCurrentProd({
                    ...currentProd,
                    name: e.target.value,
                  });
                }}
                id="outlined-besic"
                label="PROD STAGES NAME"
                variant="outlined"
                fullWidth
              />
            </Grid>
          </Grid>

          {/* CheckBoxes */}

          <Grid container spacing={1} className={classes.TextField}>
            <Grid item xs={2}>
              <FormControlLabel
                control={
                  <Checkbox
                    disabled={
                      viewMode !== viewModesNEW && viewMode !== viewModesEDIT
                    }
                    checked={currentProd.machine}
                    onChange={(e) => {
                      setCurrentProd({
                        ...currentProd,
                        machine: e.target.checked,
                      });
                    }}
                    onClick={(e) => {}}
                    name="checked"
                    color="primary"
                  />
                }
                label="MACHINE"
              />
            </Grid>
            <Grid item xs={10}>
              <TextField
                size="small"
                value={currentProd.machinename}
                disabled={
                  viewMode !== viewModesNEW && viewMode !== viewModesEDIT
                }
                onChange={(e) => {
                  setCurrentProd({
                    ...currentProd,
                    machinename: e.target.value,
                  });
                }}
                id="outlined-besic"
                label="NAME"
                variant="outlined"
                fullWidth
              />
            </Grid>
          </Grid>
          <Grid container spacing={1} className={classes.TextField}>
            <Grid item xs={2}>
              <FormControlLabel
                control={
                  <Checkbox
                    disabled={
                      viewMode !== viewModesNEW && viewMode !== viewModesEDIT
                    }
                    checked={currentProd.human}
                    onChange={(e) => {
                      setCurrentProd({
                        ...currentProd,
                        human: e.target.checked,
                      });
                    }}
                    onClick={(e) => {}}
                    name="checked"
                    color="primary"
                  />
                }
                label="HUMAN"
              />
            </Grid>
            <Grid item xs={4}>
              <TextField
                size="small"
                value={currentProd.humancount}
                type="number"
                disabled={
                  viewMode !== viewModesNEW && viewMode !== viewModesEDIT
                }
                onChange={(e) => {
                  setCurrentProd({
                    ...currentProd,
                    humancount: parseInt(e.target.value),
                  });
                }}
                id="outlined-besic"
                label="PERSON"
                variant="outlined"
                fullWidth
              />
            </Grid>
          </Grid>
        </Paper>

        {/* costhours */}

        <Grid container spacing={1} className={classes.Table}>
          <Grid item xs={6}>
            <TextField
              size="small"
              value={currentProd.costhours}
              type="number"
              disabled={viewMode !== viewModesNEW && viewMode !== viewModesEDIT}
              onChange={(e) => {
                setCurrentProd({
                  ...currentProd,
                  costhours: parseFloat(e.target.value),
                });
              }}
              id="outlined-besic"
              label="COST/HOURS"
              variant="outlined"
              fullWidth
            />
          </Grid>
        </Grid>
      </Grid>
    </>
  );
};

export default ProdStages;
