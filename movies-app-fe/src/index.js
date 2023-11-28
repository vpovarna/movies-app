import React from 'react';
import ReactDOM from 'react-dom/client';
import { RouterProvider, createBrowserRouter } from 'react-router-dom'
import App from './App';
import ErrorPage from './components/ErrorPage';
import Home from './components/Home';
import Movies from './components/Movies';
import Movie from './components/Movie';
import Genres from './components/Genres';
import EditMovies from './components/EditMovies';
import ManageCatalog from './components/ManageCatalog';
import GraphQL from './components/GraphQL';
import Login from './components/Login';

const router  = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPage />,
    children: [
      {index: true, element: <Home />},
      {
        path: "/movies", 
        element: <Movies />
      },
      {
        path: "/movies/:id", 
        element: <Movie />
      },
      {
        path: "/genres", 
        element: <Genres />
      },
      {
        path: "/admin/movies/0", 
        element: <EditMovies />
      },
      {
        path: "/manage-catalog", 
        element: <ManageCatalog />
      },
      {
        path: "/graphql", 
        element: <GraphQL />
      },
      {
        path: "/login", 
        element: <Login />
      }
    ]
  }
])

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
