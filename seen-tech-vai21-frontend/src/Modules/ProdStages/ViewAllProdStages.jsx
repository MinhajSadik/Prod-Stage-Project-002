import { Button, Grid } from "@material-ui/core";
import Paper from "@material-ui/core/Paper";
import { makeStyles } from "@material-ui/core/styles";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableContainer from "@material-ui/core/TableContainer";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import Axios from "axios";
import React, { useEffect, useState } from "react";
import { IoIosGitCompare, IoIosPaper } from "react-icons/io";
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
  ActiveColor: {
    color: "green",
  },
  InactiveColor: {
    color: "red",
  },
  MarginLeft: {
    marginLeft: theme.spacing(4),
  },
  MarginBottom: {
    marginBottom: theme.spacing(1),
  },
}));

const ViewAllProdStages = ({
  changeView,
  viewModesTOBEEIDT,
  currentProd,
  setCurrentProd,
  setViewMode,
}) => {
  const classes = useStyles();

  //my states
  const [allProdStages, setAllProdStages] = useState([]);

  const handleGetAllProdStage = async () => {
    await Axios({
      method: "GET",
      url: `${apiBaseUrl}/prodstages/get_all`,
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((res) => {
        console.log(res.data);
        res.data ? setAllProdStages(res.data.result) : setAllProdStages([]);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  const toggleStatus = async (id, status) => {
    setCurrentProd({ ...currentProd, status: status });
    const newValue = {
      ...currentProd,
      status: status,
    };

    await Axios({
      method: "PUT",
      url: `${apiBaseUrl}/prodstages/set_status/${id}/${status}`,
      headers: {
        "Content-Type": "application/json",
      },
      data: JSON.stringify(newValue),
    })
      .then((res) => {
        console.log("Status Modifing Successfully", res);
      })
      .catch((err) => {
        console.log("Error Occuard", err);
      });
  };

  //Load All Data From Api
  useEffect(() => {
    handleGetAllProdStage();
  }, [currentProd]);

  // Edited Function
  const handleChange = (prods) => {
    setCurrentProd({ ...prods, _id: prods._id });
    console.log(prods);
    setViewMode(viewModesTOBEEIDT);
  };

  return (
    <>
      <Grid container>
        <Grid item xs={12}>
          <TableContainer component={Paper}>
            <Table
              aria-label="simple-table"
              className={classes.Table}
              size="small"
            >
              <TableHead>
                <TableRow>
                  <TableCell>
                    <b>ACTIONS</b>
                  </TableCell>
                  <TableCell>
                    <b>NAME</b>
                  </TableCell>
                  <TableCell>
                    <b>MACHINE</b>
                  </TableCell>
                  <TableCell>
                    <b>HUMAN</b>
                  </TableCell>
                  <TableCell>
                    <b>COST/HOURS</b>
                  </TableCell>
                  <TableCell>
                    <b>STATUS</b>
                  </TableCell>
                </TableRow>
              </TableHead>

              <TableBody>
                {allProdStages.map((prodstages) => (
                  <TableRow key={prodstages._id}>
                    <TableCell component="th" scope="row">
                      <Button
                        className={classes.Button}
                        variant="contained"
                        color="primary"
                        onClick={() => {
                          handleChange(prodstages);
                          changeView(0);
                        }}
                      >
                        <IoIosPaper />
                      </Button>

                      <Button
                        className={classes.Button}
                        variant="contained"
                        color="primary"
                        value={currentProd.status}
                        onClick={(e) => {
                          prodstages.status
                            ? toggleStatus(prodstages._id, "inactive")
                            : toggleStatus(prodstages._id, "active");
                        }}
                      >
                        <IoIosGitCompare />
                      </Button>
                    </TableCell>
                    <TableCell>{prodstages.name}</TableCell>
                    <TableCell>
                      {prodstages.machine ? prodstages.machinename : "N/A"}
                    </TableCell>
                    <TableCell>
                      {prodstages.human ? prodstages.humancount : "N/A"}
                    </TableCell>
                    <TableCell>{prodstages.costhours}</TableCell>

                    {prodstages.status ? (
                      <TableCell className={classes.ActiveColor}>
                        Active
                      </TableCell>
                    ) : (
                      <TableCell className={classes.InactiveColor}>
                        Inactive
                      </TableCell>
                    )}
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
        </Grid>
      </Grid>
    </>
  );
};

export default ViewAllProdStages;
