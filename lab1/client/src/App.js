import { useEffect, useState } from "react";

const UNKNOWN = "UNKNOWN"
const UNAVAILABLE = "UNAVAILABLE"

function App() {
	const [alexServiceState, alexServiceSetState] = useState(UNKNOWN);
	const [titovServiceState, titovServiceSetState] = useState(UNKNOWN);
	const [maxServiceState, maxServiceSetState] = useState(UNKNOWN);

	useEffect(() => {
		fetch('/api/max-service/')
			.then((response) => {
				return response.text();
			})
			.then((data) => {
				maxServiceSetState(data);
			}).catch(() => {
				maxServiceSetState(UNAVAILABLE);
			});
	}, []);

	useEffect(() => {
		fetch('/api/titov-service/')
			.then((response) => {
				return response.text();
			})
			.then((data) => {
				titovServiceSetState(data);
			}).catch(() => {
				titovServiceSetState(UNAVAILABLE);
			});
	}, []);

	useEffect(() => {
		fetch('/api/golang-service/')
			.then((response) => {
				return response.text();
			})
			.then((data) => {
				alexServiceSetState(data);
			}).catch(() => {
				alexServiceSetState(UNAVAILABLE);
			});
	}, []);

	return (
		<div>
			golang-service status: {alexServiceState}<br />
			titov-service status: {titovServiceState}<br />
			max-service status: {maxServiceState}
		</div>
	);
}

export default App;
