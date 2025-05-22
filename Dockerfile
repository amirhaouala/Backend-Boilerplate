FROM alpine:3

RUN addgroup --system --gid 1001 golang
RUN adduser --system --uid 1001 chi

COPY --chown=chi:golang back /usr/bin
COPY --chown=chi:golang .env /usr/bin

USER chi