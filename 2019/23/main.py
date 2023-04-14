from time import time, sleep
from threading import Thread
from sys import path

path.append("../05")
from computer import Computer

start = time()

with open("input.txt", encoding="utf8") as file:
    memory = [int(x) for x in file.read().strip().split(",")]


class NIC(Computer):
    def set_input(self):
        if len(self.inputs) == 0:
            self.inputs = [-1]
            self.attempt += 1
            sleep(0.001)
        else:
            self.attempt = 0

        self.set(self.memory_address(1), self.inputs.pop(0))
        self.pc += 2


nics, nat_packets = [], []


def terminate():
    for nic in nics:
        nic.terminated = True


def send_message(address, message):
    if address == 255:
        nat_packets.append(message)
        return

    destination_nic = nics[address]
    destination_nic.inputs.extend(message)


def run_nat():
    last_y_sent = 0

    while True:
        all_nics_idle = all(nic.attempt > 5 for nic in nics)

        if all_nics_idle:
            nics[0].inputs.extend(nat_packets[-1])

            y_sent = nat_packets[-1][1]
            if y_sent == last_y_sent:
                terminate()
                break

            last_y_sent = y_sent

        sleep(0.01)


def run_nic(address):
    nic = nics[address]

    while not nic.terminated:
        dest_address = nic.run()
        message = [nic.run(), nic.run()]

        if not nic.terminated:
            send_message(dest_address, message)


threads = []
for address in range(50):
    nics.append(NIC(memory, [address]))
    threads.append(Thread(target=lambda x=address: run_nic(x)))

threads.append(Thread(target=run_nat))

[t.start() for t in threads]
[t.join() for t in threads]

time_elapsed = round(time() - start, 5)

print(
    f"""The Y value of the first packet sent to address 255 is {nat_packets[0][1]}.
The Y value of the first duplicated NAT packet is {nat_packets[-1][1]}.
Solution generated in {time_elapsed}s."""
)
