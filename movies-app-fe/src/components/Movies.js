import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

const Movies = () => {
    const [movies, setMovies] = useState([]);


    useEffect(() => {
        let moviesList = [
            {
                id: 1,
                title: "Highlander",
                release_data: "1986-03-07",
                runtime: 116,
                mpaa_rating: "R",
                description: "Some long description"
            },
            {
                id: 2,
                title: "Riders of the lost Ark",
                release_data: "1981-06-12",
                runtime: 115,
                mpaa_rating: "PG-13",
                description: "Some long description"
            },
        ];

        setMovies(moviesList)
    }, []);

    return (
        <div>
            <h2>Movies</h2>
            <hr />
            <table className="table table-striped table-hover">
                <thread>
                    <tr>
                        <th>Movie</th>
                        <th>Release Data</th>
                        <th>Rating</th>
                    </tr>
                </thread>
                <tbody>
                    {movies.map((m) => (
                        <tr key={m.id}>
                            <td>
                                <Link to={`/movies/${m.id}`}>{m.title}</Link>
                            </td>
                            <td>{m.release_data}</td>
                            <td>{m.mpaa_rating}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    )
}

export default Movies;
