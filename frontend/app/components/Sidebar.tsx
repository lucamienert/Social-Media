import { Link } from "react-router";

// src/components/Sidebar.js
function Sidebar() {
  return (
    <div className="d-flex">
      {/* Sidebar */}
      <div className="bg-dark text-white" style={{ width: '250px', minHeight: '100vh' }}>
        <div className="p-3">
          <h2>Sidebar</h2>
        </div>
        <ul className="nav flex-column">
          <li className="nav-item">
            <Link to="/" className="nav-link text-white">Home</Link>
          </li>
          <li className="nav-item">
            <Link to="/about" className="nav-link text-white">About</Link>
          </li>
          <li className="nav-item">
            <Link to="/login" className="nav-link text-white">Login</Link>
          </li>
          <li className="nav-item">
            <Link to="/register" className="nav-link text-white">Register</Link>
          </li>
          <li className="nav-item">
            <Link to="/user" className="nav-link text-white">User</Link>
          </li>
        </ul>
      </div>

      {/* Main Content */}
      <div className="flex-grow-1">
      </div>
    </div>
  );
}

export default Sidebar;
