import React, { useEffect, useState } from 'react';
import logo from './logo.svg';
import './App.css';
import axios from 'axios';
import { Table } from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';

type UserType = {
  ID: number;
  NAME: string;
};

const App: React.FC<{}> = () => {
  const [users, setUsers] = useState<UserType[]>([]);
  useEffect(() => {
    getGolangData();
  }, []);
  const getGolangData = async () => {
    try {
      await axios.get('/api/users').then((res) => {
        console.log(res.data);
        setUsers(res.data);
      });
    } catch (error) {
      console.log(error);
    }
  };
  console.log(users);
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <Table striped bordered hover>
          <thead>
            <tr>
              <th>ID</th>
              <th>Name</th>
            </tr>
          </thead>
          <tbody>
            {users.map((user, index) => {
              return (
                <tr key={index}>
                  <td>{user.ID}</td>
                  <td>{user.NAME}</td>
                </tr>
              );
            })}
          </tbody>
        </Table>
      </header>
    </div>
  );
};

export default App;
