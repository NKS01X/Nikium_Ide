FROM golang:1.22-alpine

# Create non-root user
RUN adduser -D sandbox

WORKDIR /workspace
RUN chown -R sandbox:sandbox /workspace

USER sandbox

# Keep container alive and run the command given by worker
ENTRYPOINT ["go", "run", "main.go"]
