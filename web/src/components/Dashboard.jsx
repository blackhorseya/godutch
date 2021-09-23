import React from 'react';
import {Grid} from '@material-ui/core';

const Dashboard = () => {
  return (
      <React.Fragment>
        <Grid item xs={false} sm={2}/>
        <Grid item xs={12} sm={8} container direction="column" spacing={1}>
          <Grid item>
            Hello world
          </Grid>
        </Grid>
        <Grid item xs={false} sm={2}/>
      </React.Fragment>
  );
};

export default Dashboard;