import React, {useEffect} from 'react';
import {
  Button,
  Card,
  CardContent,
  CircularProgress,
  Grid,
  makeStyles,
  Typography,
} from '@material-ui/core';
import ExitToAppIcon from '@material-ui/icons/ExitToApp';
import {useDispatch, useSelector} from 'react-redux';
import {logoutRequest} from '../redux/ducks/userSlice';
import {useHistory} from 'react-router-dom';

const useStyles = makeStyles({
  root: {
    minWidth: 275,
  },
  bullet: {
    display: 'inline-block',
    margin: '0 2px',
    transform: 'scale(0.8)',
  },
  title: {
    fontSize: 14,
  },
  pos: {
    marginBottom: 12,
  },
  logoutButton: {
    marginTop: 20,
  },
});

const Profile = (props) => {
  const classes = useStyles();
  const dispatch = useDispatch();
  const history = useHistory();
  const {profile, loading} = useSelector(state => state.user);

  useEffect(() => {
    if (profile === undefined) {
      history.push('/');
    }
  }, [history, profile]);

  function handleLogout() {
    dispatch(logoutRequest({token: localStorage.getItem('token') || ''}));
  }

  return (
      <Grid container style={{margin: 20}}>
        <Grid item xs={false} sm={2}/>
        <Grid item xs={12} sm={8}>
          {loading ? <CircularProgress/> : profile &&
              <Card className={classes.root}>
                <CardContent>
                  <Typography className={classes.title} color="textSecondary"
                              gutterBottom>
                    Name
                  </Typography>
                  <Typography variant="h5" component="h2">
                    {profile.name}
                  </Typography>
                  <Typography className={classes.title} color="textSecondary">
                    Email
                  </Typography>
                  <Typography variant="h5" component="h2">
                    {profile.email}
                  </Typography>
                  <Button className={classes.logoutButton}
                          onClick={() => handleLogout()}
                          startIcon={<ExitToAppIcon/>} color="secondary">
                    Logout
                  </Button>
                </CardContent>
              </Card>
          }
        </Grid>
        <Grid item xs={false} sm={2}/>
      </Grid>
  );
};

export default Profile;