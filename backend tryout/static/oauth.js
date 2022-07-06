const {
  colors,
  CssBaseline,
  ThemeProvider,
  Typography,
  Container,
  Paper,
  makeStyles,
  createMuiTheme,
  Box,
  CircularProgress,
  SvgIcon,
  Link,
  Grid,
} = MaterialUI;

const { useEffect, useState } = React;

const theme = createMuiTheme({
	  palette: {
		      primary: {
			            main: "#556cd6",
			          },
		      secondary: {
			            main: "#19857b",
			          },
		      error: {
			            main: colors.red.A400,
			          },
		      background: {
			            default: "#fff",
			          },
		    },
});

const useStyles = makeStyles((theme) => ({
	  root: {
		      margin: theme.spacing(6, 0, 3),
		    },
	  lightBulb: {
		      verticalAlign: "middle",
		      marginRight: theme.spacing(1),
		    },
}));

const App = () => {
	  const [timer, setTimer] = useState(0);
	  const [displayText, setDisplayText] = useState("Authenticating...");

	  useEffect(() => {
		      if (timer > 2) {
			            setDisplayText("Halo, " + name);
			            setTimeout(() => {
					            window.location.href = "/dashboard";
					          }, 1000);
			          }
		      setTimeout(() => {
			            setTimer(timer + 1);
			          }, 500);
		    }, [timer]);
	return (
		    <div
		      style={{
			              width: "auto",
				              margin: "auto",
				              paddingRight: "auto",
				              paddingLeft: "auto",
				            }}
		    >
		      <Grid container>
		        <Grid item xs={12}>
		          <div
		            style={{
				                  height: "400",
					                  width: "360px",
					                  margin: "auto",
					                  position: "relative",
					                }}
		          >
		            <img
		              style={{ position: "absolute", margin: "15px" }}
		              src="https://studybuddy.id/static/img/Logo.png"
		              height="330px"
		              width="330px"
		            />
		            {displayText == "Authenticating..." ? (
				                  <CircularProgress
				                    thickness={1.5}
				                    style={{ position: "absolute" }}
				                    color="#dddddd"
				                    size={360}
				                  />
				                ) : null}
		          </div>
		        </Grid>
		      </Grid>
		<Grid container>
		        <Grid item xs={12}>
		          <Paper elevation={0} style={{ textAlign: "center" }}>
		            <Typography
		              variant="h4"
		              component="h1"
		              textAlign="center"
		              style={{ marginTop: "380px" }}
		            >
		              {displayText}
		            </Typography>
		          </Paper>
		        </Grid>
		      </Grid>
		    </div>
		  );
};
ReactDOM.render(
	  <ThemeProvider theme={theme}>
	    {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
	    <CssBaseline />
	    <App />
	  </ThemeProvider>,
)
