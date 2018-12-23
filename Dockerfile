FROM scratch

EXPOSE 80

CMD ["/redirect"]

ADD redirect /
