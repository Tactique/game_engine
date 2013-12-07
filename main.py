#!/usr/bin/env python

from engine import network

def main():
    network.listen('0.0.0.0', 5269)

if __name__ == '__main__':
    main()
