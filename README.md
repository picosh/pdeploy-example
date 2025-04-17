# pdeploy Example Repository

This repository demonstrates how to deploy applications using the `pdeploy` system. It includes:

1.  A simple Go web server ([`main.go`](./main.go)) that echoes back the incoming HTTP request details.
2.  A standard `kennethreitz/httpbin` service for testing HTTP requests.

Both services are defined in the [`docker-compose.yml`](./docker-compose.yml) file and are configured for discovery by a Traefik reverse proxy running on the `pdeploy` host.

## Deployment with pdeploy

`pdeploy` simplifies deployment onto a Linux host running Docker and Traefik. It leverages Git for deployment triggers.

1.  **Add the pdeploy remote:**
    Replace `<repo_name>` with your desired repository name on the pdeploy server (e.g., `pdeploy-test`).
    ```bash
    git remote add pdeploy pdeploy:<repo_name>.git
    ```

2.  **Push to deploy:**
    Pushing to the `main` branch triggers the deployment process.
    ```bash
    git push pdeploy main
    ```

**What happens during deployment:**

*   The push is received by the `pdeploy` server.
*   A Git repository is created or updated at `/repos/<repo_name>.git`.
*   A post-receive hook clones the repository to `/projects/<repo_name>`.
*   `docker compose up --build --remove-orphans -d` is executed within `/projects/<repo_name>`, building the `echo` service image (using [`Dockerfile`](./Dockerfile)) and starting both services.
*   Build and deployment logs are streamed back to your Git client.

## Accessing Services

Once deployed, the services are accessible via the Traefik proxy based on the labels in [`docker-compose.yml`](./docker-compose.yml):

*   **Echo Service:** `http://echo-antonio.pdeploy.mik.qa`
*   **HTTPBin Service:** `http://http-antonio.pdeploy.mik.qa`

*(Note: Replace `antonio` with your username if applicable, based on the `pdeploy` server's routing configuration: `*-user.pdeploy.mik.qa`)*

## Managing the Deployment

You can manage the running containers via SSH:

*   **SSH into the project directory:**
    ```bash
    ssh <repo_name>@pdeploy
    ```
    This logs you directly into the `/projects/<repo_name>` directory on the server.

*   **Run docker-compose commands remotely:**
    ```bash
    # View running services
    ssh <repo_name>@pdeploy ps

    # Stop services
    ssh <repo_name>@pdeploy down

    # Start services
    ssh <repo_name>@pdeploy up -d

    # View logs
    ssh <repo_name>@pdeploy logs -f
    ```

*   **SSH into the root projects directory:**
    ```bash
    ssh pdeploy
    ```
    This logs you into `/projects`.