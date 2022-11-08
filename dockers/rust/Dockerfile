FROM rust:1.65.0


WORKDIR /usr/src/superapp
COPY . .
RUN cargo build --release

EXPOSE 3000
EXPOSE 3200

CMD ["/usr/src/target/release/superapp"]

# as builder
# #FROM rust:1.65.0-alpine as builder
# #RUN apk add --no-cache musl-dev
# #RUN mkdir opt/ ; mkdir opt/superapp
# #WORKDIR /opt/superupp

# #RUN USER=root cargo new --bin suoeraap
# WORKDIR /opt/superapp
# COPY ./ ./
# RUN cargo build --release
# #RUN rm ./src/*.rs
# #RUN rm ./target/release/deps/superupp*

# RUN ls -a /opt/superapp/
# RUN ls -a /opt/superapp/target
# RUN ls -a /opt/superapp/target/release/
# RUN ls -a /opt/superapp/target/release/superapp

# ADD ./src ./src
# RUN cargo build --release

# FROM scratch
# WORKDIR /opt/superupp
# COPY --from=builder /opt/superapp/target/release/superapp .
# COPY ./data ./data

# EXPOSE 3000
# EXPOSE 3200
# CMD ["/opt/superapp/superapp"]

